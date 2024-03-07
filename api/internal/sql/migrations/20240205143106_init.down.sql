-- Drop indexes from 'role_permissions' table
DROP INDEX IF EXISTS idx_role_permissions_role_id;
DROP INDEX IF EXISTS idx_role_permissions_permission_id;

-- Drop 'role_permissions' table
DROP TABLE IF EXISTS role_permissions;

-- Drop indexes from 'permissions' table
DROP INDEX IF EXISTS idx_permissions_internal_name;

-- Drop 'permissions' table
DROP TABLE IF EXISTS permissions;

-- Drop indexes from 'users_validation' table
DROP INDEX IF EXISTS idx_users_validation_user_id;
DROP INDEX IF EXISTS idx_users_validation_hash;
DROP INDEX IF EXISTS idx_users_validation_used;
DROP INDEX IF EXISTS idx_users_validation_created_at ;

-- Drop 'users_validation' table
DROP TABLE IF EXISTS users_validation;

-- Drop indexes from 'users' table
DROP INDEX IF EXISTS idx_users_role_id;
DROP INDEX IF EXISTS idx_users_name;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_active;

-- Drop 'users' table
DROP TABLE IF EXISTS users;

-- Drop indexes from 'roles' table
DROP INDEX IF EXISTS idx_roles_internal_name;

-- Drop 'roles' table
DROP TABLE IF EXISTS roles;

-- Drop indexes from 'avatars' table
DROP INDEX IF EXISTS idx_avatars_svg;

-- Drop 'avatars' table
DROP TABLE IF EXISTS avatars;