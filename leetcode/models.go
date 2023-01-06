package leetcode

type UserStatus struct {
	Username        string `json:"username"`
	UserSlug        string `json:"userSlug"`
	RealName        string `json:"realName"`
	Avatar          string `json:"avatar"`
	ActiveSessionId int    `json:"activeSessionId"`
	IsSignedIn      bool   `json:"isSignedIn"`
	IsPremium       bool   `json:"isPremium"`
}

type InterpretSolutionResult struct {
	InterpretExpectedId string `json:"interpret_expected_id"`
	InterpretId         string `json:"interpret_id"`
	TestCase            string `json:"test_case"`
}

type SubmitCheckResult struct {
	CodeOutput        string  `json:"code_output"`
	CompareResult     string  `json:"compare_result"`
	ElapsedTime       int     `json:"elapsed_time"`
	ExpectedOutput    string  `json:"expected_output"`
	FastSubmit        bool    `json:"fast_submit"`
	Finished          bool    `json:"finished"`
	Lang              string  `json:"lang"`
	LastTestcase      string  `json:"last_testcase"`
	Memory            int     `json:"memory"`
	MemoryPercentile  float64 `json:"memory_percentile"`
	PrettyLang        string  `json:"pretty_lang"`
	QuestionId        string  `json:"question_id"`
	RunSuccess        bool    `json:"run_success"`
	RuntimePercentile float64 `json:"runtime_percentile"`
	State             string  `json:"state"`
	StatusCode        int     `json:"status_code"`
	StatusMemory      string  `json:"status_memory"`
	StatusMsg         string  `json:"status_msg"`
	StatusRuntime     string  `json:"status_runtime"`
	StdOutput         string  `json:"std_output"`
	SubmissionId      string  `json:"submission_id"`
	TaskFinishTime    int     `json:"task_finish_time"`
	TaskName          string  `json:"task_name"`
	TotalCorrect      int     `json:"total_correct"`
	TotalTestcases    int     `json:"total_testcases"`
	CompileError      string  `json:"compile_error"`
	FullCompileError  string  `json:"full_compile_error"`
	FullRuntimeError  string  `json:"full_runtime_error"`
}

type TestCheckResult struct {
	State                  string   `json:"state"` // STARTED, SUCCESS
	StatusCode             int      `json:"status_code"`
	StatusMsg              string   `json:"status_msg"`         // Accepted, Wrong Answer, Time Limit Exceeded, Memory Limit Exceeded, Runtime Error, Compile Error, Output Limit Exceeded, Unknown Error
	Memory                 int      `json:"memory"`             // 内存消耗 in bytes
	StatusMemory           string   `json:"status_memory"`      // 内存消耗
	MemoryPercentile       float64  `json:"memory_percentile"`  // 内存消耗击败百分比
	StatusRuntime          string   `json:"status_runtime"`     // 执行用时
	RuntimePercentile      float64  `json:"runtime_percentile"` // 用时击败百分比
	Lang                   string   `json:"lang"`
	PrettyLang             string   `json:"pretty_lang"`
	CodeAnswer             []string `json:"code_answer"`   // return values of our code
	CompileError           string   `json:"compile_error"` //
	FullCompileError       string   `json:"full_compile_error"`
	FullRuntimeError       string   `json:"full_runtime_error"`
	CompareResult          string   `json:"compare_result"`  // "111", 1 means correct, 0 means wrong
	CorrectAnswer          bool     `json:"correct_answer"`  // true means all passed
	CodeOutput             []string `json:"code_output"`     // output to stdout of our code
	StdOutputList          []string `json:"std_output_list"` // list of output to stdout, same as code_output
	TaskName               string   `json:"task_name"`
	TotalCorrect           int      `json:"total_correct"`   // 通过测试用例
	TotalTestcases         int      `json:"total_testcases"` // 总测试用例
	ElapsedTime            int      `json:"elapsed_time"`
	TaskFinishTime         int      `json:"task_finish_time"`
	RunSuccess             bool     `json:"run_success"` // true if run success
	FastSubmit             bool     `json:"fast_submit"`
	Finished               bool     `json:"finished"`
	ExpectedOutput         string   `json:"expected_output"`
	ExpectedCodeAnswer     []string `json:"expected_code_answer"`
	ExpectedCodeOutput     []string `json:"expected_code_output"`
	ExpectedElapsedTime    int      `json:"expected_elapsed_time"`
	ExpectedLang           string   `json:"expected_lang"`
	ExpectedMemory         int      `json:"expected_memory"`
	ExpectedRunSuccess     bool     `json:"expected_run_success"`
	ExpectedStatusCode     int      `json:"expected_status_code"`
	ExpectedStatusRuntime  string   `json:"expected_status_runtime"`
	ExpectedStdOutputList  []string `json:"expected_std_output_list"`
	ExpectedTaskFinishTime int      `json:"expected_task_finish_time"`
	ExpectedTaskName       string   `json:"expected_task_name"`
}
