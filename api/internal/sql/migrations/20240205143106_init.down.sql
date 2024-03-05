-- Drop indexes from 'role_permissions' table
DROP INDEX IF EXISTS idx_role_permissions_role_id;
DROP INDEX IF EXISTS idx_role_permissions_permission_id;

-- Drop 'role_permissions' table
DROP TABLE IF EXISTS role_permissions;
-- Drop indexes from 'users' table
DROP INDEX IF EXISTS idx_users_role_id;
DROP INDEX IF EXISTS idx_users_name;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_active;

-- Drop 'users' table
DROP TABLE users;

-- Drop indexes from 'permissions' table
DROP INDEX IF EXISTS idx_permissions_internal_name;

-- Drop 'permissions' table
DROP TABLE permissions;

-- Drop indexes from 'roles' table
DROP INDEX IF EXISTS idx_roles_internal_name;

-- Drop 'roles' table
DROP TABLE IF EXISTS roles;

-- Drop indexes from 'avatars' table
DROP INDEX IF EXISTS idx_avatars_svg;

-- Drop 'avatars' table
DROP TABLE IF EXISTS avatars;