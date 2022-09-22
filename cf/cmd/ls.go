package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls PREFIX",
	Short: "list stacks",
	Long:  `List cloudformation stacks matching PREFIX, or all if none given.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			listStacks("")
		default:
			listStacks(args[0])
		}
	},
}

// listStacks returns all stack summaries matching prefix
func listStacks(prefix string) {
	for _, s := range stacks(prefix) {
		fmt.Println(*s.StackName)
	}
}

// existingStackStatuses returns all non-deleted statuses
func existingStackStatuses() []types.StackStatus {
	return []types.StackStatus{
		types.StackStatusCreateComplete,
		types.StackStatusCreateFailed,
		types.StackStatusCreateInProgress,
		types.StackStatusDeleteFailed,
		types.StackStatusDeleteInProgress,
		types.StackStatusReviewInProgress,
		types.StackStatusRollbackComplete,
		types.StackStatusRollbackFailed,
		types.StackStatusRollbackInProgress,
		types.StackStatusUpdateComplete,
		types.StackStatusUpdateCompleteCleanupInProgress,
		types.StackStatusUpdateInProgress,
		types.StackStatusUpdateRollbackComplete,
		types.StackStatusUpdateRollbackCompleteCleanupInProgress,
		types.StackStatusUpdateRollbackFailed,
		types.StackStatusUpdateRollbackInProgress,
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
