// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"context"
	"errors"
	"github.com/cloudwego-contrib/cwgo-open-analysis/model"
	"github.com/cloudwego-contrib/cwgo-open-analysis/util"

	"gorm.io/gorm"
)

func CreateGroup(ctx context.Context, db *gorm.DB, group *model.Group) error {
	return db.WithContext(ctx).Create(group).Error
}

func UpdateGroup(ctx context.Context, db *gorm.DB, group *model.Group) error {
	var currentGroup model.Group
	if err := db.WithContext(ctx).Where("name = ?", group.Name).First(&currentGroup).Error; err != nil {
		return err
	}
	currentGroup.IssueCount = group.IssueCount
	currentGroup.PullRequestCount = group.PullRequestCount
	currentGroup.StarCount = group.StarCount
	currentGroup.ForkCount = group.ForkCount
	currentGroup.ContributorCount = group.ContributorCount
	if err := db.WithContext(ctx).Save(&currentGroup).Error; err != nil {
		return err
	}
	return nil
}

func CreateOrganization(ctx context.Context, db *gorm.DB, org *model.Organization) error {
	return db.WithContext(ctx).Create(org).Error
}

func UpdateOrganization(ctx context.Context, db *gorm.DB, org *model.Organization) error {
	var currentOrg model.Organization
	if err := db.WithContext(ctx).Where("node_id = ?", org.NodeID).First(&currentOrg).Error; err != nil {
		return err
	}
	currentOrg.IssueCount = org.IssueCount
	currentOrg.PullRequestCount = org.PullRequestCount
	currentOrg.StarCount = org.StarCount
	currentOrg.ForkCount = org.ForkCount
	currentOrg.ContributorCount = org.ContributorCount
	if err := db.WithContext(ctx).Save(&currentOrg).Error; err != nil {
		return err
	}
	return nil
}

func CreateRepository(ctx context.Context, db *gorm.DB, repo *model.Repository) error {
	return db.WithContext(ctx).Create(repo).Error
}

func QueryRepositoryNodeID(ctx context.Context, db *gorm.DB, owner, name string) (string, error) {
	var repo model.Repository
	err := db.WithContext(ctx).Where(model.Repository{
		Owner: owner,
		Name:  name,
	}).First(&repo).Error
	return repo.NodeID, err
}

func DeleteRepository(ctx context.Context, db *gorm.DB, nodeID string) error {
	return db.WithContext(ctx).Where("node_id = ?", nodeID).Delete(&model.Repository{}).Error
}

func CreateGroupsOrganizations(ctx context.Context, db *gorm.DB, join *model.GroupsOrganizations) error {
	return db.WithContext(ctx).Create(join).Error
}

func CreateGroupsRepositories(ctx context.Context, db *gorm.DB, join *model.GroupsRepositories) error {
	return db.WithContext(ctx).Create(join).Error
}

func CreateIssues(ctx context.Context, db *gorm.DB, issues []*model.Issue) error {
	if util.IsEmptySlice(issues) {
		return nil
	}
	return db.WithContext(ctx).Create(issues).Error
}

func IssueExist(ctx context.Context, db *gorm.DB, nodeID string) (bool, error) {
	var issue model.Issue
	if err := db.WithContext(ctx).Where("node_id = ?", nodeID).First(&issue).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func UpdateIssue(ctx context.Context, db *gorm.DB, issue *model.Issue) error {
	var currentIssue model.Issue
	if err := db.WithContext(ctx).Where("node_id = ?", issue.NodeID).First(&currentIssue).Error; err != nil {
		return err
	}
	currentIssue.State = issue.State
	currentIssue.IssueClosedAt = issue.IssueClosedAt
	if err := db.WithContext(ctx).Save(&currentIssue).Error; err != nil {
		return err
	}
	return nil
}

func DeleteIssues(ctx context.Context, db *gorm.DB, repoNodeID string) error {
	return db.WithContext(ctx).Where("repo_node_id = ?", repoNodeID).Delete(&model.Issue{}).Error
}

func CreatePullRequests(ctx context.Context, db *gorm.DB, prs []*model.PullRequest) error {
	if util.IsEmptySlice(prs) {
		return nil
	}
	return db.WithContext(ctx).Create(prs).Error
}

func UpdatePullRequest(ctx context.Context, db *gorm.DB, pr *model.PullRequest) error {
	var currentPR model.PullRequest
	if err := db.WithContext(ctx).Where("node_id = ?", pr.NodeID).First(&currentPR).Error; err != nil {
		return err
	}
	currentPR.State = pr.State
	currentPR.PRMergedAt = pr.PRMergedAt
	currentPR.PRClosedAt = pr.PRClosedAt
	if err := db.WithContext(ctx).Save(&currentPR).Error; err != nil {
		return err
	}
	return nil
}

func DeletePullRequests(ctx context.Context, db *gorm.DB, repoNodeID string) error {
	return db.WithContext(ctx).Where("repo_node_id = ?", repoNodeID).Delete(&model.PullRequest{}).Error
}

func QueryOPENPullRequests(ctx context.Context, db *gorm.DB, repoNodeID string) ([]model.PullRequest, error) {
	var prs []model.PullRequest
	err := db.WithContext(ctx).Where("state = ? AND repo_node_id = ?", "OPEN", repoNodeID).Find(&prs).Error
	return prs, err
}

func CreateIssueAssignees(ctx context.Context, db *gorm.DB, assignees []*model.IssueAssignees) error {
	if util.IsEmptySlice(assignees) {
		return nil
	}
	return db.WithContext(ctx).Create(assignees).Error
}

func IssueAssigneesExist(ctx context.Context, db *gorm.DB, nodeID string) (bool, error) {
	var assignees model.IssueAssignees
	if err := db.WithContext(ctx).Where("issue_node_id = ?", nodeID).First(&assignees).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func UpdateIssueAssignees(ctx context.Context, db *gorm.DB, issueNodeID string, assignees []*model.IssueAssignees) error {
	if util.IsEmptySlice(assignees) {
		return nil
	}
	var currentAssignees []*model.IssueAssignees
	if err := db.WithContext(ctx).Where("issue_node_id = ?", issueNodeID).Find(&currentAssignees).Error; err != nil {
		return err
	}
	var s1 []model.IssueAssignees
	var s2 []model.IssueAssignees
	for _, assignee := range currentAssignees {
		s1 = append(s1, model.IssueAssignees{
			IssueNodeID:    assignee.IssueNodeID,
			IssueNumber:    assignee.IssueNumber,
			IssueURL:       assignee.IssueURL,
			IssueRepoName:  assignee.IssueRepoName,
			AssigneeNodeID: assignee.AssigneeNodeID,
			AssigneeLogin:  assignee.AssigneeLogin,
		})
	}
	for _, assignee := range assignees {
		s2 = append(s2, model.IssueAssignees{
			IssueNodeID:    assignee.IssueNodeID,
			IssueNumber:    assignee.IssueNumber,
			IssueURL:       assignee.IssueURL,
			IssueRepoName:  assignee.IssueRepoName,
			AssigneeNodeID: assignee.AssigneeNodeID,
			AssigneeLogin:  assignee.AssigneeLogin,
		})
	}
	more, less := util.CompareSlices(s1, s2)
	if !util.IsEmptySlice(more) {
		if err := db.WithContext(ctx).Create(more).Error; err != nil {
			return err
		}
	}
	for _, e := range less {
		if err := db.WithContext(ctx).Where("id = ?", e.ID).Delete(&model.IssueAssignees{}).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeleteIssueAssigneesByIssue(ctx context.Context, db *gorm.DB, issueNodeID string) error {
	return db.WithContext(ctx).Where("issue_node_id = ?", issueNodeID).Delete(&model.IssueAssignees{}).Error
}

func DeleteIssueAssigneesByRepo(ctx context.Context, db *gorm.DB, nameWithOwner string) error {
	return db.WithContext(ctx).Where("issue_repo_name = ?", nameWithOwner).Delete(&model.IssueAssignees{}).Error
}

func PullRequestAssigneesExist(ctx context.Context, db *gorm.DB, nodeID string) (bool, error) {
	var assignees model.PullRequestAssignees
	if err := db.WithContext(ctx).Where("pull_request_node_id = ?", nodeID).First(&assignees).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreatePullRequestAssignees(ctx context.Context, db *gorm.DB, assignees []*model.PullRequestAssignees) error {
	if util.IsEmptySlice(assignees) {
		return nil
	}
	return db.WithContext(ctx).Create(assignees).Error
}

func UpdatePullRequestAssignees(ctx context.Context, db *gorm.DB, prNodeID string, assignees []*model.PullRequestAssignees) error {
	if util.IsEmptySlice(assignees) {
		return nil
	}
	var currentAssignees []*model.PullRequestAssignees
	if err := db.WithContext(ctx).Where("pull_request_node_id = ?", prNodeID).Find(&currentAssignees).Error; err != nil {
		return err
	}
	var s1 []model.PullRequestAssignees
	var s2 []model.PullRequestAssignees
	for _, assignee := range currentAssignees {
		s1 = append(s1, model.PullRequestAssignees{
			PullRequestNodeID:   assignee.PullRequestNodeID,
			PullRequestNumber:   assignee.PullRequestNumber,
			PullRequestURL:      assignee.PullRequestURL,
			PullRequestRepoName: assignee.PullRequestRepoName,
			AssigneeNodeID:      assignee.AssigneeNodeID,
			AssigneeLogin:       assignee.AssigneeLogin,
		})
	}
	for _, assignee := range assignees {
		s2 = append(s2, model.PullRequestAssignees{
			PullRequestNodeID:   assignee.PullRequestNodeID,
			PullRequestNumber:   assignee.PullRequestNumber,
			PullRequestURL:      assignee.PullRequestURL,
			PullRequestRepoName: assignee.PullRequestRepoName,
			AssigneeNodeID:      assignee.AssigneeNodeID,
			AssigneeLogin:       assignee.AssigneeLogin,
		})
	}
	more, less := util.CompareSlices(s1, s2)
	if !util.IsEmptySlice(more) {
		if err := db.WithContext(ctx).Create(more).Error; err != nil {
			return err
		}
	}
	for _, e := range less {
		if err := db.WithContext(ctx).Where("id = ?", e.ID).Delete(&model.PullRequestAssignees{}).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeletePullRequestAssigneesByPR(ctx context.Context, db *gorm.DB, prNodeID string) error {
	return db.WithContext(ctx).Where("pull_request_node_id = ?", prNodeID).Delete(&model.PullRequestAssignees{}).Error
}

func DeletePullRequestAssigneesByRepo(ctx context.Context, db *gorm.DB, nameWithOwner string) error {
	return db.WithContext(ctx).Where("pull_request_repo_name = ?", nameWithOwner).Delete(&model.PullRequestAssignees{}).Error
}

func CreateContributors(ctx context.Context, db *gorm.DB, cs []*model.Contributor) error {
	if util.IsEmptySlice(cs) {
		return nil
	}
	return db.WithContext(ctx).Create(cs).Error
}

func UpdateContributorCompanyAndLocation(ctx context.Context, db *gorm.DB, update func(string) string) error {
	var contributors []model.Contributor
	if err := db.WithContext(ctx).Find(&contributors).Error; err != nil {
		return err
	}
	for _, contributor := range contributors {
		contributor.Company = update(contributor.Company)
		contributor.Location = update(contributor.Location)
		if err := db.WithContext(ctx).Save(&contributor).Error; err != nil {
			return err
		}
	}
	return nil
}

// QueryContributorCountByOrg
//
// SELECT COUNT(DISTINCT c.node_id) AS contributor_count
// FROM contributors c
// INNER JOIN repositories r ON c.repo_node_id = r.node_id
// WHERE r.owner_node_id = 'orgNodeID';
func QueryContributorCountByOrg(ctx context.Context, db *gorm.DB, orgNodeID string) (int, error) {
	var contributorCount int
	if err := db.WithContext(ctx).
		Table("contributors").
		Select("COUNT(DISTINCT contributors.node_id) AS contributor_count").
		Joins("INNER JOIN repositories ON contributors.repo_node_id = repositories.node_id").
		Where("repositories.owner_node_id = ?", orgNodeID).
		Scan(&contributorCount).Error; err != nil {
		return 0, err
	}
	return contributorCount, nil
}

// QueryContributorCountByGroup
//
// SELECT COUNT(DISTINCT c.node_id) AS contributor_count
// FROM contributors c
// INNER JOIN (
//
//	SELECT DISTINCT gr.repo_node_id
//	FROM groups_repositories gr
//	INNER JOIN repositories r ON gr.repo_node_id = r.node_id
//	WHERE gr.group_name = 'groupName'
//	UNION
//	SELECT DISTINCT r.node_id
//	FROM repositories r
//	INNER JOIN groups_organizations go ON r.owner_node_id = go.org_node_id
//	WHERE go.group_name = 'groupName'
//
// ) AS repos ON c.repo_node_id = repos.repo_node_id;
func QueryContributorCountByGroup(ctx context.Context, db *gorm.DB, groupName string) (int, error) {
	var count int64

	var repos1 []string
	sq1 := db.WithContext(ctx).
		Table("groups_repositories").
		Select("groups_repositories.repo_node_id").
		Joins("INNER JOIN repositories ON groups_repositories.repo_node_id = repositories.node_id").
		Where("groups_repositories.group_name = ?", groupName)
	if err := sq1.Find(&repos1).Error; err != nil {
		return 0, err
	}

	var repos2 []string
	sq2 := db.WithContext(ctx).
		Table("repositories").
		Select("repositories.node_id").
		Joins("INNER JOIN groups_organizations ON repositories.owner_node_id = groups_organizations.org_node_id").
		Where("groups_organizations.group_name = ?", groupName)
	if err := sq2.Find(&repos2).Error; err != nil {
		return 0, err
	}

	repoNodeIDs := append(repos1, repos2...)

	if err := db.WithContext(ctx).
		Table("contributors").
		Select("contributors.node_id").
		Where("contributors.repo_node_id IN ?", repoNodeIDs).
		Distinct().
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func UpdateOrCreateContributors(ctx context.Context, db *gorm.DB, cs []*model.Contributor) error {
	for _, contributor := range cs {
		if err := db.WithContext(ctx).Where(model.Contributor{
			NodeID:     contributor.NodeID,
			RepoNodeID: contributor.RepoNodeID,
		}).Assign(contributor).FirstOrCreate(contributor).Error; err != nil {
			return err
		}
	}
	return nil
}

func CreateCursor(ctx context.Context, db *gorm.DB, cursor *model.Cursor) error {
	return db.WithContext(ctx).Create(cursor).Error
}

func QueryCursor(ctx context.Context, db *gorm.DB, repo string) (*model.Cursor, error) {
	cursor := &model.Cursor{}
	err := db.WithContext(ctx).Where("repo_name_with_owner = ?", repo).First(cursor).Error
	// for organization's new repository case
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return cursor, nil
	}
	return cursor, err
}

func UpdateOrCreateCursor(ctx context.Context, db *gorm.DB, cursor *model.Cursor) error {
	var currentCursor model.Cursor
	if err := db.WithContext(ctx).Where("repo_node_id = ?", cursor.RepoNodeID).First(&currentCursor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.WithContext(ctx).Create(&model.Cursor{
				RepoNodeID:        cursor.RepoNodeID,
				RepoNameWithOwner: cursor.RepoNameWithOwner,
				LastUpdate:        cursor.LastUpdate,
				EndCursor:         cursor.EndCursor,
			}).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}
	currentCursor.RepoNameWithOwner = cursor.RepoNameWithOwner
	currentCursor.LastUpdate = cursor.LastUpdate
	currentCursor.EndCursor = cursor.EndCursor
	if err := db.WithContext(ctx).Save(&currentCursor).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCursor(ctx context.Context, db *gorm.DB, repoNodeID string) error {
	return db.WithContext(ctx).Where("repo_node_id = ?", repoNodeID).Delete(&model.Cursor{}).Error
}

func QueryReposByOrg(ctx context.Context, db *gorm.DB, orgNodeID string) ([]string, error) {
	var repos []model.Repository
	if err := db.WithContext(ctx).Where("owner_node_id = ?", orgNodeID).Group("node_id").Find(&repos).Error; err != nil {
		return nil, err
	}
	var reposNameWithOwner []string
	for _, repo := range repos {
		reposNameWithOwner = append(reposNameWithOwner, util.MergeNameWithOwner(repo.Owner, repo.Name))
	}
	return reposNameWithOwner, nil
}