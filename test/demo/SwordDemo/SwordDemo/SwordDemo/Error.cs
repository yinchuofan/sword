using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace SwordDemo
{
    public struct Error
    {
        public int code;
        public string message;
    }

    class ErrorCode
    {
        public static int ErrNoError = 0;
        public static int ErrHasError = 1;
        public static int ErrException = 2;

        public static int ErrServerError = 10001;
        public static int ErrParamsType = 10002;
        public static int ErrForm = 10003; 
        public static int ErrNotFound = 10004;
        public static int ErrGetCookieFailed = 10005;
        public static int ErrMessage = 10006;
        public static int ErrPushMessage = 10007;
        public static int ErrQuery = 10008; 
        public static int ErrFakeRequest = 10009; 
        public static int ErrPermissionDenied = 10010; 
        public static int ErrSendShortMessage = 10011; 

        public static int ErrNotLogin = 10100; 
        public static int ErrNotPassAuthenticated = 10101;
        public static int ErrRegisterFailed = 10102; 
        public static int ErrAccountExisting = 10103;
        public static int ErrLoginFailed = 10104;
        public static int ErrAuthenticateFailed = 10105;
        public static int ErrPhoneTypeErr = 10106;
        public static int ErrPhoneValidate = 10107; 
        public static int ErrModifyPasswordError = 10108; 
    }
}
