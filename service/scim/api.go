// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package scim

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewCurrentUser(client *client.DatabricksClient) CurrentUserService {
	return &CurrentUserAPI{
		client: client,
	}
}

type CurrentUserAPI struct {
	client *client.DatabricksClient
}

// Get current user info
//
// Get details about the current method caller's identity.
func (a *CurrentUserAPI) Me(ctx context.Context) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Me"
	err := a.client.Get(ctx, path, nil, &user)
	return &user, err
}

func NewGroups(client *client.DatabricksClient) GroupsService {
	return &GroupsAPI{
		client: client,
	}
}

type GroupsAPI struct {
	client *client.DatabricksClient
}

// Create a new group
//
// Creates a group in the <Workspace> with a unique name, using the supplied
// group details.
func (a *GroupsAPI) CreateGroup(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Post(ctx, path, request, &group)
	return &group, err
}

// Delete a group
//
// Deletes a group from the <Workspace>.
func (a *GroupsAPI) DeleteGroup(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a group
//
// Deletes a group from the <Workspace>.
func (a *GroupsAPI) DeleteGroupById(ctx context.Context, id string) error {
	return a.DeleteGroup(ctx, DeleteGroupRequest{
		Id: id,
	})
}

// Get group details
//
// Gets the information for a specific group in the <Workspace>.
func (a *GroupsAPI) GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Get(ctx, path, request, &group)
	return &group, err
}

// Get group details
//
// Gets the information for a specific group in the <Workspace>.
func (a *GroupsAPI) GetGroupById(ctx context.Context, id string) (*Group, error) {
	return a.GetGroup(ctx, GetGroupRequest{
		Id: id,
	})
}

// List group details
//
// Gets all details of the groups associated with the <Workspace>.
//
// Use ListGroupsAll() to get all Group instances
func (a *GroupsAPI) ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Get(ctx, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

// ListGroupsAll returns all Group instances. This method exists for consistency purposes.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GroupsAPI) ListGroupsAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	response, err := a.ListGroups(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

func (a *GroupsAPI) GroupDisplayNameToIdMap(ctx context.Context, request ListGroupsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListGroupsAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.DisplayName] = v.Id
	}
	return mapping, nil
}

func (a *GroupsAPI) GetGroupByDisplayName(ctx context.Context, name string) (*Group, error) {
	result, err := a.ListGroupsAll(ctx, ListGroupsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.DisplayName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Group named '%s' does not exist", name)
}

// Update group details
//
// Partially updates the details of a group.
func (a *GroupsAPI) PatchGroup(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Replace a group
//
// Updates the details of a group by replacing the entire group entity.
func (a *GroupsAPI) UpdateGroup(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}

func NewServicePrincipals(client *client.DatabricksClient) ServicePrincipalsService {
	return &ServicePrincipalsAPI{
		client: client,
	}
}

type ServicePrincipalsAPI struct {
	client *client.DatabricksClient
}

// Create a service principal
//
// Creates a new service principal in the <Workspace>.
func (a *ServicePrincipalsAPI) CreateServicePrincipal(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Post(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Delete a service principal
//
// Delete a single service principal in the <Workspace>.
func (a *ServicePrincipalsAPI) DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a service principal
//
// Delete a single service principal in the <Workspace>.
func (a *ServicePrincipalsAPI) DeleteServicePrincipalById(ctx context.Context, id string) error {
	return a.DeleteServicePrincipal(ctx, DeleteServicePrincipalRequest{
		Id: id,
	})
}

// Get service principal details
//
// Gets the details for a single service principal define in the <Workspace>.
func (a *ServicePrincipalsAPI) GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Get(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Get service principal details
//
// Gets the details for a single service principal define in the <Workspace>.
func (a *ServicePrincipalsAPI) GetServicePrincipalById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.GetServicePrincipal(ctx, GetServicePrincipalRequest{
		Id: id,
	})
}

// List service principals
//
// Gets the set of service principals associated with a <Workspace>.
//
// Use ListServicePrincipalsAll() to get all ServicePrincipal instances
func (a *ServicePrincipalsAPI) ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Get(ctx, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

// ListServicePrincipalsAll returns all ServicePrincipal instances. This method exists for consistency purposes.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ServicePrincipalsAPI) ListServicePrincipalsAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	response, err := a.ListServicePrincipals(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

func (a *ServicePrincipalsAPI) ServicePrincipalDisplayNameToIdMap(ctx context.Context, request ListServicePrincipalsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListServicePrincipalsAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.DisplayName] = v.Id
	}
	return mapping, nil
}

func (a *ServicePrincipalsAPI) GetServicePrincipalByDisplayName(ctx context.Context, name string) (*ServicePrincipal, error) {
	result, err := a.ListServicePrincipalsAll(ctx, ListServicePrincipalsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.DisplayName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("ServicePrincipal named '%s' does not exist", name)
}

// Update service principal details
//
// Partially updates the details of a single service principal in the
// <Workspace>.
func (a *ServicePrincipalsAPI) PatchServicePrincipal(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Replace service principal in <Workspace>
//
// Updates the details of a single service principal.
//
// This action replaces the existing service principal with the same name.
func (a *ServicePrincipalsAPI) UpdateServicePrincipal(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}

func NewUsers(client *client.DatabricksClient) UsersService {
	return &UsersAPI{
		client: client,
	}
}

type UsersAPI struct {
	client *client.DatabricksClient
}

// Create a new user
//
// Creates a new user in the <Workspace>. This new user will also be added to
// the <Workspace> account.
func (a *UsersAPI) CreateUser(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Post(ctx, path, request, &user)
	return &user, err
}

// Delete a user
//
// Deletes a user. Deleting a user from a workspace also removes objects
// associated with the user.
func (a *UsersAPI) DeleteUser(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a user
//
// Deletes a user. Deleting a user from a workspace also removes objects
// associated with the user.
func (a *UsersAPI) DeleteUserById(ctx context.Context, id string) error {
	return a.DeleteUser(ctx, DeleteUserRequest{
		Id: id,
	})
}

// Get user details
//
// Gets information for a specific user in <Workspace>.
func (a *UsersAPI) GetUser(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Get(ctx, path, request, &user)
	return &user, err
}

// Get user details
//
// Gets information for a specific user in <Workspace>.
func (a *UsersAPI) GetUserById(ctx context.Context, id string) (*User, error) {
	return a.GetUser(ctx, GetUserRequest{
		Id: id,
	})
}

// List users
//
// Gets details for all the users associated with a <Workspace>.
//
// Use ListUsersAll() to get all User instances
func (a *UsersAPI) ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Get(ctx, path, request, &listUsersResponse)
	return &listUsersResponse, err
}

// ListUsersAll returns all User instances. This method exists for consistency purposes.
//
// This method is generated by Databricks SDK Code Generator.
func (a *UsersAPI) ListUsersAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	response, err := a.ListUsers(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Resources, nil
}

func (a *UsersAPI) UserUserNameToIdMap(ctx context.Context, request ListUsersRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListUsersAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.UserName] = v.Id
	}
	return mapping, nil
}

func (a *UsersAPI) GetUserByUserName(ctx context.Context, name string) (*User, error) {
	result, err := a.ListUsersAll(ctx, ListUsersRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.UserName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("User named '%s' does not exist", name)
}

// Update user details
//
// Partially updates a user resource by applying the supplied operations on
// specific user attributes.
func (a *UsersAPI) PatchUser(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Replace a user
//
// Replaces a user's information with the data supplied in request.
func (a *UsersAPI) UpdateUser(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Put(ctx, path, request)
	return err
}
