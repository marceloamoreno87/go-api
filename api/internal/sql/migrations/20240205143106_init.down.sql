-- Drop indexes from 'role_permissions' table
DROP INDEX IF EXISTS idx_role_permissions_role_id;
DROP INDEX IF EXISTS idx_role_permissions_permission_id;

-- Drop 'role_permissions' table
DROP TABLE IF EXISTS role_permissions;

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

-- Drop indexes from 'validation_types' table
DROP INDEX IF EXISTS idx_validation_types_name;

-- Drop 'validation_types' table
DROP TABLE IF EXISTS validation_types;

-- Drop indexes from 'validation_users' table
DROP INDEX IF EXISTS idx_validation_users_user_id;
DROP INDEX IF EXISTS idx_validation_users_hash;
DROP INDEX IF EXISTS idx_validation_users_validation_type_id;

-- Drop 'validation_users' table
DROP TABLE IF EXISTS validation_users;

-- Drop indexes from 'users' table
DROP INDEX IF EXISTS idx_users_role_id;
DROP INDEX IF EXISTS idx_users_name;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_active;

-- Drop 'users' table
DROP TABLE users;