package v1

import (
	"context"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/internal/log"
)

func (o *openAPI) Roles(ctx context.Context, guildID string) (*dto.GuildRoles, error) {
	resp, err := o.request(ctx).
		SetResult(dto.GuildRoles{}).
		SetPathParam("guild_id", guildID).
		Get(getURL(rolesURI, o.sandbox))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.GuildRoles), nil
}

func (o *openAPI) PostRole(ctx context.Context, guildID string, role *dto.Role) (dto.RoleID, error) {
	if role.Color == 0 {
		role.Color = dto.DefaultColor
	}
	// openapi 上修改哪个字段，就需要传递哪个 filter
	filter := &dto.UpdateRoleFilter{
		Name:  1,
		Color: 1,
		Hoist: 1,
	}
	body := &dto.UpdateRole{
		GuildID: guildID,
		Filter:  filter,
		Update:  role,
	}
	log.Debug(body)
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetResult(dto.UpdateResult{}).
		SetBody(body).
		Post(getURL(rolesURI, o.sandbox))
	if err != nil {
		return "", err
	}

	return resp.Result().(*dto.UpdateResult).RoleID, nil
}

func (o *openAPI) PatchRole(ctx context.Context,
	guildID string, roleID dto.RoleID, role *dto.Role) (dto.RoleID, error) {
	if role.Color == 0 {
		role.Color = dto.DefaultColor
	}
	filter := &dto.UpdateRoleFilter{
		Name:  1,
		Color: 1,
		Hoist: 1,
	}
	body := &dto.UpdateRole{
		GuildID: guildID,
		Filter:  filter,
		Update:  role,
	}
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		SetResult(dto.UpdateResult{}).
		SetBody(body).
		Patch(getURL(roleURI, o.sandbox))
	if err != nil {
		return "", err
	}

	return resp.Result().(*dto.UpdateResult).RoleID, nil
}

func (o *openAPI) DeleteRole(ctx context.Context, guildID string, roleID dto.RoleID) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		Delete(getURL(roleURI, o.sandbox))
	return err
}
