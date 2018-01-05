using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.IO;
using System.Windows.Forms;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace SwordDemo
{
    public partial class SwordDemo : Form
    {
        public SwordDemo()
        {
            InitializeComponent();
        }
        
        private void SwordDemo_Load(object sender, EventArgs e)
        {
            InitMainLayout();

            RefreshLayout();

            this.tbSwordHost.Text = this.swordHost;
        }

        private void InitMainLayout()
        {
            int screenWidth = System.Windows.Forms.SystemInformation.WorkingArea.Width;
            int screenHeight = System.Windows.Forms.SystemInformation.WorkingArea.Height;
            this.WindowState = FormWindowState.Maximized;
            this.MinimumSize = new System.Drawing.Size(screenWidth * 2 / 3, screenHeight * 2 / 3);            
        }

        private void RefreshLayout()
        {
            this.plSwordHost.Location = new Point(20,20);
            this.plSwordHost.Size = new System.Drawing.Size((this.ClientSize.Width - 20 * 3) / 2, 50);
            this.lbSwordHost.Location = new Point(10, 20);
            this.tbSwordHost.Location = new Point(94,15);
            this.tbSwordHost.Width = this.plSwordHost.Width - 94 - 10;

            this.plInputArgs.Location = new Point(20, this.plSwordHost.Location.Y + this.plSwordHost.Height + 20);
            this.plInputArgs.Size = new System.Drawing.Size((this.ClientSize.Width - 20 * 3) / 2, (this.ClientSize.Height - 20 * 4 - 50) * 3 / 4);
            this.lbTaskInfo.Location = new Point(10,10);
            this.tbTaskInfo.Location = new Point(10, 30);
            this.tbTaskInfo.Size = new System.Drawing.Size(this.plInputArgs.Width - 20,this.plInputArgs.Height - 110);
            this.lbTaskID.Location = new Point(10, this.tbTaskInfo.Location.Y + this.tbTaskInfo.Height + 20);
            this.tbTaskID.Location = new Point(10, this.lbTaskID.Location.Y + this.lbTaskID.Height + 10);
            this.tbTaskID.Width = this.tbTaskInfo.Width;

            this.plAPIList.Location = new Point(20, this.plInputArgs.Location.Y + this.plInputArgs.Height + 20);
            this.plAPIList.Size = new System.Drawing.Size((this.ClientSize.Width - 20 * 3) / 2, (this.ClientSize.Height - 20 * 4 - 50) * 1 / 4);
            this.btCreateTask.Location = new Point(20,20);
            this.btStartTask.Location = new Point(this.btCreateTask.Location.X + this.btCreateTask.Width + 10, 20);
            this.btStopTask.Location = new Point(this.btStartTask.Location.X + this.btStartTask.Width + 10, 20);
            this.btDeleteTask.Location = new Point(this.btStopTask.Location.X + this.btStopTask.Width + 10, 20);
            this.btQueryTaskStatus.Location = new Point(20, this.btCreateTask.Location.Y + this.btCreateTask.Height + 20);
            this.btQueryAllTaskStatus.Location = new Point(this.btQueryTaskStatus.Location.X + this.btQueryTaskStatus.Width + 10, this.btCreateTask.Location.Y + this.btCreateTask.Height + 20);
            this.btQueryTaskCapacity.Location = new Point(this.btQueryAllTaskStatus.Location.X + this.btQueryAllTaskStatus.Width + 10, this.btCreateTask.Location.Y + this.btCreateTask.Height + 20);
            this.btCreateTask.Width = 140;
            this.btStartTask.Width = 140;
            this.btStopTask.Width = 140;
            this.btDeleteTask.Width = 140;
            this.btQueryTaskStatus.Width = 140;
            this.btQueryAllTaskStatus.Width = 140;
            this.btQueryTaskCapacity.Width = 140;

            this.plAPIResult.Location = new Point(20 * 2 + this.plInputArgs.Width, 20);
            this.plAPIResult.Size = new System.Drawing.Size((this.ClientSize.Width - 20 * 3) / 2, this.ClientSize.Height - 20 * 2);
            this.lbAPIResult.Location = new Point(10, 10);
            this.tbAPIResult.Location = new Point(10, 30);
            this.tbAPIResult.Size = new System.Drawing.Size(this.plAPIResult.Width - 20, this.plAPIResult.Height - 40);
        }

        private void SwordDemo_Resize(object sender, EventArgs e)
        {
            RefreshLayout();
        }

        private string FormatJsonString(string str)
        {
            //格式化json字符串
            JsonSerializer serializer = new JsonSerializer();
            TextReader tr = new StringReader(str);
            JsonTextReader jtr = new JsonTextReader(tr);
            object obj = serializer.Deserialize(jtr);
            if (obj != null)
            {
                StringWriter textWriter = new StringWriter();
                JsonTextWriter jsonWriter = new JsonTextWriter(textWriter)
                {
                    Formatting = Formatting.Indented,
                    Indentation = 4,
                    IndentChar = ' '
                };
                serializer.Serialize(jsonWriter, obj);
                return textWriter.ToString();
            }
            else
            {
                return str;
            }
        }

        private void btCreateTask_Click(object sender, EventArgs e)
        {
            this.swordHost = this.tbSwordHost.Text;
            if (this.swordHost == "")
            {
                MessageBox.Show("Sword host is null");
                return;
            }

            this.taskPara = this.tbTaskInfo.Text;
            if (this.taskPara == "")
            {
                MessageBox.Show("Task information is null");
                return;
            }

            string response = "";
            Error err = HttpAPI.CreateTask(this.swordHost, this.taskPara, ref response);
            if (ErrorCode.ErrNoError != err.code)
            {
                MessageBox.Show("API call error");
                return;
            }

            this.tbAPIResult.Text = FormatJsonString(response);
            
        }

        private void btStartTask_Click(object sender, EventArgs e)
        {
            this.swordHost = this.tbSwordHost.Text;
            if (this.swordHost == "")
            {
                MessageBox.Show("Sword host is null");
                return;
            }

            string taskIDListString = this.tbTaskID.Text;
            if (taskIDListString == "")
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            this.taskIDList = taskIDListString.Split(',').ToList();
            if (this.taskIDList.Count == 0)
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            string taskID = this.taskIDList[0];

            string response = "";
            Error err = HttpAPI.StartTask(this.swordHost, taskID, ref response);
            if (ErrorCode.ErrNoError != err.code)
            {
                MessageBox.Show("API call error");
                return;
            }

            this.tbAPIResult.Text = FormatJsonString(response);
        }

        private void btStopTask_Click(object sender, EventArgs e)
        {
            this.swordHost = this.tbSwordHost.Text;
            if (this.swordHost == "")
            {
                MessageBox.Show("Sword host is null");
                return;
            }

            string taskIDListString = this.tbTaskID.Text;
            if (taskIDListString == "")
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            this.taskIDList = taskIDListString.Split(',').ToList();
            if (this.taskIDList.Count == 0)
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            string taskID = this.taskIDList[0];

            string response = "";
            Error err = HttpAPI.StopTask(this.swordHost, taskID, ref response);
            if (ErrorCode.ErrNoError != err.code)
            {
                MessageBox.Show("API call error");
                return;
            }

            this.tbAPIResult.Text = FormatJsonString(response);
        }

        private void btDeleteTask_Click(object sender, EventArgs e)
        {
            this.swordHost = this.tbSwordHost.Text;
            if (this.swordHost == "")
            {
                MessageBox.Show("Sword host is null");
                return;
            }

            string taskIDListString = this.tbTaskID.Text;
            if (taskIDListString == "")
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            this.taskIDList = taskIDListString.Split(',').ToList();
            if (this.taskIDList.Count == 0)
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            string taskID = this.taskIDList[0];

            string response = "";
            Error err = HttpAPI.DeleteTask(this.swordHost, taskID, ref response);
            if (ErrorCode.ErrNoError != err.code)
            {
                MessageBox.Show("API call error");
                return;
            }

            this.tbAPIResult.Text = FormatJsonString(response);
        }

        private void btQueryTaskStatus_Click(object sender, EventArgs e)
        {
            this.swordHost = this.tbSwordHost.Text;
            if (this.swordHost == "")
            {
                MessageBox.Show("Sword host is null");
                return;
            }

            string taskIDListString = this.tbTaskID.Text;
            if (taskIDListString == "")
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            this.taskIDList = taskIDListString.Split(',').ToList();
            if (this.taskIDList.Count == 0)
            {
                MessageBox.Show("Task ID is null");
                return;
            }

            string response = "";
            Error err = HttpAPI.QueryTaskStatus(this.swordHost, this.taskIDList, ref response);
            if (ErrorCode.ErrNoError != err.code)
            {
                MessageBox.Show("API call error");
                return;
            }

            this.tbAPIResult.Text = FormatJsonString(response);
        }

        private void btQueryAllTaskStatus_Click(object sender, EventArgs e)
        {
            this.swordHost = this.tbSwordHost.Text;
            if (this.swordHost == "")
            {
                MessageBox.Show("Sword host is null");
                return;
            }

            string response = "";
            Error err = HttpAPI.QueryAllTaskStatus(this.swordHost, ref response);
            if (ErrorCode.ErrNoError != err.code)
            {
                MessageBox.Show("API call error");
                return;
            }

            this.tbAPIResult.Text = FormatJsonString(response);
        }

        private void btQueryTaskCapacity_Click(object sender, EventArgs e)
        {
            this.swordHost = this.tbSwordHost.Text;
            if (this.swordHost == "")
            {
                MessageBox.Show("Sword host is null");
                return;
            }

            string response = "";
            Error err = HttpAPI.QueryTaskCapacity(this.swordHost, ref response);
            if (ErrorCode.ErrNoError != err.code)
            {
                MessageBox.Show("API call error");
                return;
            }

            this.tbAPIResult.Text = FormatJsonString(response);
        }

        private string swordHost = "127.0.0.1:6000";
        private string taskPara = "";
        private List<string> taskIDList = new List<string>();


    }
}
