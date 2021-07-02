using System;
using System.Threading.Tasks;
using System.Net.Http;

namespace CdgrFrontend
{
    public class CodeGraderClient
    {
        string apiKey;
        string url;

        CodeGraderClient(string url, string apiKey = "")
        {
            this.apiKey = apiKey;
            this.url = url;
        }

        async Task<GradeResult> GradeCode(string testcaseName, string userCode)
        {

            var client = new HttpClient();
            var content = new StringContent(userCode);
            content.Headers.Add("testcase-name", testcaseName);

            var res = await client.PostAsync(url + "/grade", null);

            var result = await res.Content.ReadAsStringAsync();
            return new GradeResult{};
        }
    }

    public enum GradeResultStatus {
        Success,
        WrongAnswer,
        ErrCompilation,
        ErrInternal,
        RuntimeExceed,
        MemoryLimitExceed

    }

    public struct GradeResult {
        public GradeResultStatus Status;
        public string Msg;

    }

}
