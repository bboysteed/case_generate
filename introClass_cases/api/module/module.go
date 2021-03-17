package module




type Info struct {
	AllStudents []*Student `json:"all_students"`
}

type Student struct {
	StuName string `json:"stu_name"`
	JobsNames []string `json:"jobs_names"`
}

func (st *Student) AppendJob(jobName string) {
	st.JobsNames = append(st.JobsNames, jobName)
}

func (info *Info) AppendStudent(stu *Student) {
	info.AllStudents = append(info.AllStudents, stu)
}