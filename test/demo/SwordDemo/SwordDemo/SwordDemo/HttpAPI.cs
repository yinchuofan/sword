using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Net.Http;
using System.IO;
using System.Web;
using System.Net;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace SwordDemo
{
    class HttpAPI
    {
        public static Error CreateTask(string swordHost, string taskPara, ref string response)
        {
            string HTTP_URL = "http://" + swordHost + "/createTask";

            try
            {
                StringWriter sw = new StringWriter();
                JsonWriter writer = new JsonTextWriter(sw);
                writer.WriteStartObject();
                writer.WritePropertyName("taskPara");
                writer.WriteValue(taskPara);
                writer.WriteEndObject();
                writer.Flush();
                string jsonText = sw.GetStringBuilder().ToString();

                byte[] jsonBodyBytes = Encoding.UTF8.GetBytes(jsonText);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(HTTP_URL);
                request.Method = "POST";
                request.ContentType = "application/json";
                request.ContentLength = jsonBodyBytes.Length;
                Stream reqStream = request.GetRequestStream();
                reqStream.Write(jsonBodyBytes, 0, jsonBodyBytes.Length);
                reqStream.Close();

                // get the response
                WebResponse webResponse = request.GetResponse();
                Stream webStream = webResponse.GetResponseStream();
                StreamReader responseReader = new StreamReader(webStream);
                response = responseReader.ReadToEnd();
                responseReader.Close();

                Error err = new Error();
                err.code = ErrorCode.ErrNoError;
                err.message = "OK";
                return err;
            }
            catch (System.Exception ex)
            {
                Error err = new Error();
                err.code = ErrorCode.ErrException;
                err.message = ex.Message;
                return err;
            }
        }

        public static Error StartTask(string swordHost, string taskID, ref string response)
        {
            string HTTP_URL = "http://" + swordHost + "/startTask";

            try
            {
                StringWriter sw = new StringWriter();
                JsonWriter writer = new JsonTextWriter(sw);
                writer.WriteStartObject();
                writer.WritePropertyName("taskID");
                writer.WriteValue(taskID);
                writer.WriteEndObject();
                writer.Flush();
                string jsonText = sw.GetStringBuilder().ToString();

                byte[] jsonBodyBytes = Encoding.UTF8.GetBytes(jsonText);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(HTTP_URL);
                request.Method = "POST";
                request.ContentType = "application/json";
                request.ContentLength = jsonBodyBytes.Length;
                Stream reqStream = request.GetRequestStream();
                reqStream.Write(jsonBodyBytes, 0, jsonBodyBytes.Length);
                reqStream.Close();

                // get the response
                WebResponse webResponse = request.GetResponse();
                Stream webStream = webResponse.GetResponseStream();
                StreamReader responseReader = new StreamReader(webStream);
                response = responseReader.ReadToEnd();
                responseReader.Close();

                Error err = new Error();
                err.code = ErrorCode.ErrNoError;
                err.message = "OK";
                return err;
            }
            catch (System.Exception ex)
            {
                Error err = new Error();
                err.code = ErrorCode.ErrException;
                err.message = ex.Message;
                return err;
            }
        }

        public static Error StopTask(string swordHost, string taskID, ref string response)
        {
            string HTTP_URL = "http://" + swordHost + "/stopTask";

            try
            {
                StringWriter sw = new StringWriter();
                JsonWriter writer = new JsonTextWriter(sw);
                writer.WriteStartObject();
                writer.WritePropertyName("taskID");
                writer.WriteValue(taskID);
                writer.WriteEndObject();
                writer.Flush();
                string jsonText = sw.GetStringBuilder().ToString();

                byte[] jsonBodyBytes = Encoding.UTF8.GetBytes(jsonText);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(HTTP_URL);
                request.Method = "POST";
                request.ContentType = "application/json";
                request.ContentLength = jsonBodyBytes.Length;
                Stream reqStream = request.GetRequestStream();
                reqStream.Write(jsonBodyBytes, 0, jsonBodyBytes.Length);
                reqStream.Close();

                // get the response
                WebResponse webResponse = request.GetResponse();
                Stream webStream = webResponse.GetResponseStream();
                StreamReader responseReader = new StreamReader(webStream);
                response = responseReader.ReadToEnd();
                responseReader.Close();

                Error err = new Error();
                err.code = ErrorCode.ErrNoError;
                err.message = "OK";
                return err;
            }
            catch (System.Exception ex)
            {
                Error err = new Error();
                err.code = ErrorCode.ErrException;
                err.message = ex.Message;
                return err;
            }
        }

        public static Error DeleteTask(string swordHost, string taskID, ref string response)
        {
            string HTTP_URL = "http://" + swordHost + "/deleteTask";

            try
            {
                StringWriter sw = new StringWriter();
                JsonWriter writer = new JsonTextWriter(sw);
                writer.WriteStartObject();
                writer.WritePropertyName("taskID");
                writer.WriteValue(taskID);
                writer.WriteEndObject();
                writer.Flush();
                string jsonText = sw.GetStringBuilder().ToString();

                byte[] jsonBodyBytes = Encoding.UTF8.GetBytes(jsonText);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(HTTP_URL);
                request.Method = "POST";
                request.ContentType = "application/json";
                request.ContentLength = jsonBodyBytes.Length;
                Stream reqStream = request.GetRequestStream();
                reqStream.Write(jsonBodyBytes, 0, jsonBodyBytes.Length);
                reqStream.Close();

                // get the response
                WebResponse webResponse = request.GetResponse();
                Stream webStream = webResponse.GetResponseStream();
                StreamReader responseReader = new StreamReader(webStream);
                response = responseReader.ReadToEnd();
                responseReader.Close();

                Error err = new Error();
                err.code = ErrorCode.ErrNoError;
                err.message = "OK";
                return err;
            }
            catch (System.Exception ex)
            {
                Error err = new Error();
                err.code = ErrorCode.ErrException;
                err.message = ex.Message;
                return err;
            }
        }

        public static Error QueryTaskStatus(string swordHost, List<string> taskIDList, ref string response)
        {
            string HTTP_URL = "http://" + swordHost + "/queryTaskStatus";
            string para = "";
            foreach (string taskID in taskIDList)
            {
                para = para + "taskID=" + taskID + "&";
            }
            if (para != "")
            {
                para = para.Substring(0, para.Length - 1);
            }            

            HTTP_URL = HTTP_URL + "?" + para;

            try
            {
                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(HTTP_URL);
                request.Method = "GET";

                // get the response
                WebResponse webResponse = request.GetResponse();
                Stream webStream = webResponse.GetResponseStream();
                StreamReader responseReader = new StreamReader(webStream);
                response = responseReader.ReadToEnd();
                responseReader.Close();

                Error err = new Error();
                err.code = ErrorCode.ErrNoError;
                err.message = "OK";
                return err;
            }
            catch (System.Exception ex)
            {
                Error err = new Error();
                err.code = ErrorCode.ErrException;
                err.message = ex.Message;
                return err;
            }
        }

        public static Error QueryAllTaskStatus(string swordHost, ref string response)
        {
            string HTTP_URL = "http://" + swordHost + "/queryAllTaskStatus";

            try
            {
                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(HTTP_URL);
                request.Method = "GET";

                // get the response
                WebResponse webResponse = request.GetResponse();
                Stream webStream = webResponse.GetResponseStream();
                StreamReader responseReader = new StreamReader(webStream);
                response = responseReader.ReadToEnd();
                responseReader.Close();

                Error err = new Error();
                err.code = ErrorCode.ErrNoError;
                err.message = "OK";
                return err;
            }
            catch (System.Exception ex)
            {
                Error err = new Error();
                err.code = ErrorCode.ErrException;
                err.message = ex.Message;
                return err;
            }
        }

        public static Error QueryTaskCapacity(string swordHost, ref string response)
        {
            string HTTP_URL = "http://" + swordHost + "/queryTaskCapacity";

            try
            {
                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(HTTP_URL);
                request.Method = "GET";

                // get the response
                WebResponse webResponse = request.GetResponse();
                Stream webStream = webResponse.GetResponseStream();
                StreamReader responseReader = new StreamReader(webStream);
                response = responseReader.ReadToEnd();
                responseReader.Close();

                Error err = new Error();
                err.code = ErrorCode.ErrNoError;
                err.message = "OK";
                return err;
            }
            catch (System.Exception ex)
            {
                Error err = new Error();
                err.code = ErrorCode.ErrException;
                err.message = ex.Message;
                return err;
            }
        }
    }
}
