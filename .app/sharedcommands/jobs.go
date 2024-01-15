package sharedcommands

var JobsCmd 

func init(){
	JobsCmd := &cobra.Command{
	Use:   "jobs",
	Short: "Jobs",
	
}
RootCmd.AddCommand(JobsCmd)
}