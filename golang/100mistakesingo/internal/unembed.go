package internal

type UnEmbed struct {
	I int
}

type IMple struct {
	i UnEmbed // unexported
	S string
}
