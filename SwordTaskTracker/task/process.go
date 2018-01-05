package task

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"sword/SwordJobTracker/log"
)

func StartTaskProcess(taskID, taskPara string) (processID int, err error) {
	taskProcessPath := TaskPath

	//	switch runtime.GOOS {
	//	case "linux":
	//		taskProcessPath = TaskPath + "task"
	//	case "windows":
	//		taskProcessPath = TaskPath + "task.exe"
	//	default:
	//		taskProcessPath = TaskPath + "task"
	//	}

	cmd := exec.Command(taskProcessPath, taskPara)

	if err = cmd.Start(); err != nil {
		log.Root.Error("Start task process error. TaskID: %v, TaskPath: %v", taskID, taskProcessPath)
		return -1, err
	}

	monitorProcess(taskID, cmd)
	return cmd.Process.Pid, nil
}

func monitorProcess(taskID string, cmd *exec.Cmd) {
	switch runtime.GOOS {
	case "linux":
		go monitorProcessInLinux(taskID, cmd)
	case "windows":
		go monitorProcessInWindows(taskID, cmd)
	default:
		go monitorProcessInLinux(taskID, cmd)
	}
}

func monitorProcessInWindows(taskID string, cmd *exec.Cmd) {
	if err := cmd.Wait(); err != nil {
		//Process exited abnormally for some sutiation below
		//1) exited with non-zero
		//2) coredump as divide by zero and so on
		//3) be killed by host process
		//If process be killed by host process, will report finished, otherwise report exception
		status := cmd.ProcessState.Sys().(syscall.WaitStatus)
		if status.Exited() && status.ExitStatus() == 1 {
			//Process be killed by host process
			//Report task finished to master node
			ReportTaskFinished(taskID)
		} else {
			//Exited with non-zero or coredump
			//Report task finished to master node
			ReportTaskException(taskID)
		}

	} else {
		//Process exited normally and exit with zero
		//Report task finished to master node
		ReportTaskFinished(taskID)
	}

	//Task process exited, remove task from compute node
	removeTask(taskID)
}

func monitorProcessInLinux(taskID string, cmd *exec.Cmd) {
	if err := cmd.Wait(); err != nil {
		//Process exited abnormally for some sutiation below
		//1) exited with non-zero
		//2) coredump as divide by zero and so on
		//3) be killed by host process
		//If process be killed by host process, will report finished, otherwise report exception
		status := cmd.ProcessState.Sys().(syscall.WaitStatus)
		if status.Signaled() && status.Signal() == syscall.SIGKILL {
			//Process be killed by host process
			log.Root.Info("Process receive kill signal. TaskID: %v, ProcessID: %v", taskID, cmd.Process.Pid)

			//Report task finished to master node
			ReportTaskFinished(taskID)
		} else {
			//Exited with non-zero or coredump
			log.Root.Info("Process exit with exception. TaskID: %v, ProcessID: %v", taskID, cmd.Process.Pid)

			//Report task finished to master node
			ReportTaskException(taskID)
		}

	} else {
		//Process exited normally and exit with zero
		log.Root.Info("Process exit normally. TaskID: %v, ProcessID: %v", taskID, cmd.Process.Pid)

		//Report task finished to master node
		ReportTaskFinished(taskID)
	}

	//Task process exited, remove task from compute node
	removeTask(taskID)
}

func StopTaskProcess(processID int) (err error) {
	p, err := os.FindProcess(processID)
	if err != nil {
		log.Root.Error("Find process error. ProcessID: %v", processID)
		return err
	}

	if runtime.GOOS == "linux" {
		if err = p.Signal(syscall.SIGINT); err != nil {
			log.Root.Error("Interrupt process error. ProcessID: %v", processID)
			return err
		}
	} else {
		if err = p.Kill(); err != nil {
			log.Root.Error("Kill process error. ProcessID: %v", processID)
			return err
		}
	}

	if err = p.Release(); err != nil {
		log.Root.Error("Release process error. ProcessID: %v", processID)
		return err
	}

	return nil
}
