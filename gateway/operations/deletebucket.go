package operations

import (
	"treeverse-lake/gateway/errors"
	"treeverse-lake/gateway/permissions"
)

type DeleteBucket struct{}

func (controller *DeleteBucket) GetArn() string {
	return "arn:treeverse:repos:::*"
}

func (controller *DeleteBucket) GetPermission() string {
	return permissions.PermissionManageRepos
}

func (controller *DeleteBucket) Handle(o *RepoOperation) {
	err := o.Index.DeleteRepo(o.ClientId, o.Repo)
	if err != nil {
		o.EncodeError(errors.Codes.ToAPIErr(errors.ErrInternalError))
		return
	}
}