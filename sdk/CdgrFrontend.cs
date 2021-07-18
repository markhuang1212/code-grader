using System;
using System.Threading.Tasks;
using System.Net.Http;

namespace CdgrFrontend
{

    public class GradeRequest
    {
        string TestCaseName;
        string Usercode;
    }

    public class CodeGraderClient
    {
        string apiKey;
        string endpoint;

        CodeGraderClient(string apiKey = "",string endpoint = "")
        {
            this.apiKey = apiKey;
            this.endpoint = endpoint;
        }

        async Task<string> GradeUserCode(GradeRequest gr)
        {
            return "1234";
        }

        async Task<GradeResult> GetGradeResult(string id)
        {
            var ret = new GradeResult();
            return ret;
        }

        async Task<string[]> GetTestCases()
        {
            return new string[] { "TestCase1", "TestCase2" };
        }

    }

    public enum GradeResultStatus
    {
        GradeResultSuccess,
        GradeResultWrongAnswer,
        GradeResultExecutionError,
        GradeResultInternalError,
        GradeResultCompilationError,
        GradeResultTimeLimitExceed,
        GradeResultMemoryExceed,
    }

    public struct GradeResult
    {
        public GradeResultStatus Status;
        public long Duration;
        public string Msg;
    }

}
