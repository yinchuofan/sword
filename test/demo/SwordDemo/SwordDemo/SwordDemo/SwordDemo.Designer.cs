namespace SwordDemo
{
    partial class SwordDemo
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            this.plInputArgs = new System.Windows.Forms.Panel();
            this.lbTaskInfo = new System.Windows.Forms.Label();
            this.lbTaskID = new System.Windows.Forms.Label();
            this.tbTaskID = new System.Windows.Forms.TextBox();
            this.plAPIList = new System.Windows.Forms.Panel();
            this.btCreateTask = new System.Windows.Forms.Button();
            this.btStartTask = new System.Windows.Forms.Button();
            this.btStopTask = new System.Windows.Forms.Button();
            this.btDeleteTask = new System.Windows.Forms.Button();
            this.btQueryTaskStatus = new System.Windows.Forms.Button();
            this.btQueryAllTaskStatus = new System.Windows.Forms.Button();
            this.btQueryTaskCapacity = new System.Windows.Forms.Button();
            this.tbTaskInfo = new System.Windows.Forms.TextBox();
            this.plAPIResult = new System.Windows.Forms.Panel();
            this.lbAPIResult = new System.Windows.Forms.Label();
            this.tbAPIResult = new System.Windows.Forms.TextBox();
            this.plSwordHost = new System.Windows.Forms.Panel();
            this.lbSwordHost = new System.Windows.Forms.Label();
            this.tbSwordHost = new System.Windows.Forms.TextBox();
            this.plInputArgs.SuspendLayout();
            this.plAPIList.SuspendLayout();
            this.plAPIResult.SuspendLayout();
            this.plSwordHost.SuspendLayout();
            this.SuspendLayout();
            // 
            // plInputArgs
            // 
            this.plInputArgs.BackColor = System.Drawing.SystemColors.Control;
            this.plInputArgs.Controls.Add(this.tbTaskInfo);
            this.plInputArgs.Controls.Add(this.tbTaskID);
            this.plInputArgs.Controls.Add(this.lbTaskID);
            this.plInputArgs.Controls.Add(this.lbTaskInfo);
            this.plInputArgs.Location = new System.Drawing.Point(3, 62);
            this.plInputArgs.Name = "plInputArgs";
            this.plInputArgs.Size = new System.Drawing.Size(433, 249);
            this.plInputArgs.TabIndex = 0;
            // 
            // lbTaskInfo
            // 
            this.lbTaskInfo.AutoSize = true;
            this.lbTaskInfo.Location = new System.Drawing.Point(10, 14);
            this.lbTaskInfo.Name = "lbTaskInfo";
            this.lbTaskInfo.Size = new System.Drawing.Size(107, 12);
            this.lbTaskInfo.TabIndex = 0;
            this.lbTaskInfo.Text = "Task Information:";
            // 
            // lbTaskID
            // 
            this.lbTaskID.AutoSize = true;
            this.lbTaskID.Location = new System.Drawing.Point(12, 191);
            this.lbTaskID.Name = "lbTaskID";
            this.lbTaskID.Size = new System.Drawing.Size(53, 12);
            this.lbTaskID.TabIndex = 2;
            this.lbTaskID.Text = "Task ID:";
            // 
            // tbTaskID
            // 
            this.tbTaskID.Location = new System.Drawing.Point(9, 215);
            this.tbTaskID.Name = "tbTaskID";
            this.tbTaskID.Size = new System.Drawing.Size(406, 21);
            this.tbTaskID.TabIndex = 3;
            // 
            // plAPIList
            // 
            this.plAPIList.BackColor = System.Drawing.SystemColors.Control;
            this.plAPIList.Controls.Add(this.btDeleteTask);
            this.plAPIList.Controls.Add(this.btQueryTaskCapacity);
            this.plAPIList.Controls.Add(this.btQueryAllTaskStatus);
            this.plAPIList.Controls.Add(this.btStopTask);
            this.plAPIList.Controls.Add(this.btQueryTaskStatus);
            this.plAPIList.Controls.Add(this.btStartTask);
            this.plAPIList.Controls.Add(this.btCreateTask);
            this.plAPIList.Location = new System.Drawing.Point(4, 323);
            this.plAPIList.Name = "plAPIList";
            this.plAPIList.Size = new System.Drawing.Size(432, 98);
            this.plAPIList.TabIndex = 1;
            // 
            // btCreateTask
            // 
            this.btCreateTask.BackColor = System.Drawing.Color.SeaGreen;
            this.btCreateTask.FlatAppearance.BorderSize = 0;
            this.btCreateTask.FlatAppearance.MouseDownBackColor = System.Drawing.Color.Brown;
            this.btCreateTask.FlatAppearance.MouseOverBackColor = System.Drawing.Color.Brown;
            this.btCreateTask.FlatStyle = System.Windows.Forms.FlatStyle.Flat;
            this.btCreateTask.ForeColor = System.Drawing.Color.White;
            this.btCreateTask.Location = new System.Drawing.Point(13, 16);
            this.btCreateTask.Name = "btCreateTask";
            this.btCreateTask.Size = new System.Drawing.Size(97, 23);
            this.btCreateTask.TabIndex = 0;
            this.btCreateTask.Text = "Create Task";
            this.btCreateTask.UseVisualStyleBackColor = false;
            this.btCreateTask.Click += new System.EventHandler(this.btCreateTask_Click);
            // 
            // btStartTask
            // 
            this.btStartTask.BackColor = System.Drawing.Color.SeaGreen;
            this.btStartTask.FlatAppearance.BorderSize = 0;
            this.btStartTask.FlatAppearance.MouseDownBackColor = System.Drawing.Color.Brown;
            this.btStartTask.FlatAppearance.MouseOverBackColor = System.Drawing.Color.Brown;
            this.btStartTask.FlatStyle = System.Windows.Forms.FlatStyle.Flat;
            this.btStartTask.ForeColor = System.Drawing.Color.White;
            this.btStartTask.Location = new System.Drawing.Point(116, 16);
            this.btStartTask.Name = "btStartTask";
            this.btStartTask.Size = new System.Drawing.Size(97, 23);
            this.btStartTask.TabIndex = 0;
            this.btStartTask.Text = "Start Task";
            this.btStartTask.UseVisualStyleBackColor = false;
            this.btStartTask.Click += new System.EventHandler(this.btStartTask_Click);
            // 
            // btStopTask
            // 
            this.btStopTask.BackColor = System.Drawing.Color.SeaGreen;
            this.btStopTask.FlatAppearance.BorderSize = 0;
            this.btStopTask.FlatAppearance.MouseDownBackColor = System.Drawing.Color.Brown;
            this.btStopTask.FlatAppearance.MouseOverBackColor = System.Drawing.Color.Brown;
            this.btStopTask.FlatStyle = System.Windows.Forms.FlatStyle.Flat;
            this.btStopTask.ForeColor = System.Drawing.Color.White;
            this.btStopTask.Location = new System.Drawing.Point(219, 16);
            this.btStopTask.Name = "btStopTask";
            this.btStopTask.Size = new System.Drawing.Size(97, 23);
            this.btStopTask.TabIndex = 0;
            this.btStopTask.Text = "Stop Task";
            this.btStopTask.UseVisualStyleBackColor = false;
            this.btStopTask.Click += new System.EventHandler(this.btStopTask_Click);
            // 
            // btDeleteTask
            // 
            this.btDeleteTask.BackColor = System.Drawing.Color.SeaGreen;
            this.btDeleteTask.FlatAppearance.BorderSize = 0;
            this.btDeleteTask.FlatAppearance.MouseDownBackColor = System.Drawing.Color.Brown;
            this.btDeleteTask.FlatAppearance.MouseOverBackColor = System.Drawing.Color.Brown;
            this.btDeleteTask.FlatStyle = System.Windows.Forms.FlatStyle.Flat;
            this.btDeleteTask.ForeColor = System.Drawing.Color.White;
            this.btDeleteTask.Location = new System.Drawing.Point(322, 16);
            this.btDeleteTask.Name = "btDeleteTask";
            this.btDeleteTask.Size = new System.Drawing.Size(97, 23);
            this.btDeleteTask.TabIndex = 0;
            this.btDeleteTask.Text = "Delete Task";
            this.btDeleteTask.UseVisualStyleBackColor = false;
            this.btDeleteTask.Click += new System.EventHandler(this.btDeleteTask_Click);
            // 
            // btQueryTaskStatus
            // 
            this.btQueryTaskStatus.BackColor = System.Drawing.Color.SeaGreen;
            this.btQueryTaskStatus.FlatAppearance.BorderSize = 0;
            this.btQueryTaskStatus.FlatAppearance.MouseDownBackColor = System.Drawing.Color.Brown;
            this.btQueryTaskStatus.FlatAppearance.MouseOverBackColor = System.Drawing.Color.Brown;
            this.btQueryTaskStatus.FlatStyle = System.Windows.Forms.FlatStyle.Flat;
            this.btQueryTaskStatus.ForeColor = System.Drawing.Color.White;
            this.btQueryTaskStatus.Location = new System.Drawing.Point(13, 61);
            this.btQueryTaskStatus.Name = "btQueryTaskStatus";
            this.btQueryTaskStatus.Size = new System.Drawing.Size(122, 23);
            this.btQueryTaskStatus.TabIndex = 0;
            this.btQueryTaskStatus.Text = "Query Task Status";
            this.btQueryTaskStatus.UseVisualStyleBackColor = false;
            this.btQueryTaskStatus.Click += new System.EventHandler(this.btQueryTaskStatus_Click);
            // 
            // btQueryAllTaskStatus
            // 
            this.btQueryAllTaskStatus.BackColor = System.Drawing.Color.SeaGreen;
            this.btQueryAllTaskStatus.FlatAppearance.BorderSize = 0;
            this.btQueryAllTaskStatus.FlatAppearance.MouseDownBackColor = System.Drawing.Color.Brown;
            this.btQueryAllTaskStatus.FlatAppearance.MouseOverBackColor = System.Drawing.Color.Brown;
            this.btQueryAllTaskStatus.FlatStyle = System.Windows.Forms.FlatStyle.Flat;
            this.btQueryAllTaskStatus.ForeColor = System.Drawing.Color.White;
            this.btQueryAllTaskStatus.Location = new System.Drawing.Point(141, 61);
            this.btQueryAllTaskStatus.Name = "btQueryAllTaskStatus";
            this.btQueryAllTaskStatus.Size = new System.Drawing.Size(141, 23);
            this.btQueryAllTaskStatus.TabIndex = 0;
            this.btQueryAllTaskStatus.Text = "Query All Task Status";
            this.btQueryAllTaskStatus.UseVisualStyleBackColor = false;
            this.btQueryAllTaskStatus.Click += new System.EventHandler(this.btQueryAllTaskStatus_Click);
            // 
            // btQueryTaskCapacity
            // 
            this.btQueryTaskCapacity.BackColor = System.Drawing.Color.SeaGreen;
            this.btQueryTaskCapacity.FlatAppearance.BorderSize = 0;
            this.btQueryTaskCapacity.FlatAppearance.MouseDownBackColor = System.Drawing.Color.Brown;
            this.btQueryTaskCapacity.FlatAppearance.MouseOverBackColor = System.Drawing.Color.Brown;
            this.btQueryTaskCapacity.FlatStyle = System.Windows.Forms.FlatStyle.Flat;
            this.btQueryTaskCapacity.ForeColor = System.Drawing.Color.White;
            this.btQueryTaskCapacity.Location = new System.Drawing.Point(288, 61);
            this.btQueryTaskCapacity.Name = "btQueryTaskCapacity";
            this.btQueryTaskCapacity.Size = new System.Drawing.Size(131, 23);
            this.btQueryTaskCapacity.TabIndex = 0;
            this.btQueryTaskCapacity.Text = "Query Task Capacity";
            this.btQueryTaskCapacity.UseVisualStyleBackColor = false;
            this.btQueryTaskCapacity.Click += new System.EventHandler(this.btQueryTaskCapacity_Click);
            // 
            // tbTaskInfo
            // 
            this.tbTaskInfo.Font = new System.Drawing.Font("宋体", 10.5F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.tbTaskInfo.Location = new System.Drawing.Point(12, 40);
            this.tbTaskInfo.Multiline = true;
            this.tbTaskInfo.Name = "tbTaskInfo";
            this.tbTaskInfo.Size = new System.Drawing.Size(406, 134);
            this.tbTaskInfo.TabIndex = 4;
            // 
            // plAPIResult
            // 
            this.plAPIResult.BackColor = System.Drawing.SystemColors.Control;
            this.plAPIResult.Controls.Add(this.tbAPIResult);
            this.plAPIResult.Controls.Add(this.lbAPIResult);
            this.plAPIResult.Location = new System.Drawing.Point(443, 4);
            this.plAPIResult.Name = "plAPIResult";
            this.plAPIResult.Size = new System.Drawing.Size(408, 418);
            this.plAPIResult.TabIndex = 2;
            // 
            // lbAPIResult
            // 
            this.lbAPIResult.AutoSize = true;
            this.lbAPIResult.Location = new System.Drawing.Point(14, 9);
            this.lbAPIResult.Name = "lbAPIResult";
            this.lbAPIResult.Size = new System.Drawing.Size(71, 12);
            this.lbAPIResult.TabIndex = 0;
            this.lbAPIResult.Text = "API Result:";
            // 
            // tbAPIResult
            // 
            this.tbAPIResult.Font = new System.Drawing.Font("宋体", 10.5F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.tbAPIResult.Location = new System.Drawing.Point(16, 30);
            this.tbAPIResult.Multiline = true;
            this.tbAPIResult.Name = "tbAPIResult";
            this.tbAPIResult.Size = new System.Drawing.Size(382, 375);
            this.tbAPIResult.TabIndex = 4;
            // 
            // plSwordHost
            // 
            this.plSwordHost.BackColor = System.Drawing.SystemColors.Control;
            this.plSwordHost.Controls.Add(this.tbSwordHost);
            this.plSwordHost.Controls.Add(this.lbSwordHost);
            this.plSwordHost.Location = new System.Drawing.Point(4, 4);
            this.plSwordHost.Name = "plSwordHost";
            this.plSwordHost.Size = new System.Drawing.Size(432, 50);
            this.plSwordHost.TabIndex = 3;
            // 
            // lbSwordHost
            // 
            this.lbSwordHost.AutoSize = true;
            this.lbSwordHost.Location = new System.Drawing.Point(9, 20);
            this.lbSwordHost.Name = "lbSwordHost";
            this.lbSwordHost.Size = new System.Drawing.Size(71, 12);
            this.lbSwordHost.TabIndex = 0;
            this.lbSwordHost.Text = "Sword Host:";
            // 
            // tbSwordHost
            // 
            this.tbSwordHost.Font = new System.Drawing.Font("宋体", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.tbSwordHost.Location = new System.Drawing.Point(94, 15);
            this.tbSwordHost.Name = "tbSwordHost";
            this.tbSwordHost.Size = new System.Drawing.Size(321, 21);
            this.tbSwordHost.TabIndex = 1;
            // 
            // SwordDemo
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.BackColor = System.Drawing.Color.Silver;
            this.ClientSize = new System.Drawing.Size(853, 434);
            this.Controls.Add(this.plSwordHost);
            this.Controls.Add(this.plAPIResult);
            this.Controls.Add(this.plAPIList);
            this.Controls.Add(this.plInputArgs);
            this.Name = "SwordDemo";
            this.ShowIcon = false;
            this.Text = "SwordDemo";
            this.Load += new System.EventHandler(this.SwordDemo_Load);
            this.Resize += new System.EventHandler(this.SwordDemo_Resize);
            this.plInputArgs.ResumeLayout(false);
            this.plInputArgs.PerformLayout();
            this.plAPIList.ResumeLayout(false);
            this.plAPIResult.ResumeLayout(false);
            this.plAPIResult.PerformLayout();
            this.plSwordHost.ResumeLayout(false);
            this.plSwordHost.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.Panel plInputArgs;
        private System.Windows.Forms.TextBox tbTaskID;
        private System.Windows.Forms.Label lbTaskID;
        private System.Windows.Forms.Label lbTaskInfo;
        private System.Windows.Forms.Panel plAPIList;
        private System.Windows.Forms.Button btDeleteTask;
        private System.Windows.Forms.Button btQueryTaskCapacity;
        private System.Windows.Forms.Button btQueryAllTaskStatus;
        private System.Windows.Forms.Button btStopTask;
        private System.Windows.Forms.Button btQueryTaskStatus;
        private System.Windows.Forms.Button btStartTask;
        private System.Windows.Forms.Button btCreateTask;
        private System.Windows.Forms.TextBox tbTaskInfo;
        private System.Windows.Forms.Panel plAPIResult;
        private System.Windows.Forms.Label lbAPIResult;
        private System.Windows.Forms.TextBox tbAPIResult;
        private System.Windows.Forms.Panel plSwordHost;
        private System.Windows.Forms.TextBox tbSwordHost;
        private System.Windows.Forms.Label lbSwordHost;
    }
}

