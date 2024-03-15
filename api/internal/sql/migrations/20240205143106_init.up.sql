-- Create the 'roles' table
CREATE TABLE IF NOT EXISTS roles (
  id SERIAL PRIMARY KEY, -- Role id
  name VARCHAR(255) UNIQUE NOT NULL, -- Role name
  internal_name VARCHAR(255) UNIQUE NOT NULL, -- Role internal name
  description TEXT NOT NULL, -- Role description
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Creation timestamp
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL -- Update timestamp
);
CREATE INDEX idx_roles_internal_name ON roles(internal_name);

-- Insert roles
INSERT INTO roles (name, internal_name, description) VALUES ('Administrator', 'admin', 'Administrator user');
INSERT INTO roles (name, internal_name, description) VALUES ('User', 'user', 'Common user');
INSERT INTO roles (name, internal_name, description) VALUES ('Analyst', 'analyst', 'Analyst user');

-- Create the 'permissions' table
CREATE TABLE IF NOT EXISTS permissions (
  id SERIAL PRIMARY KEY, -- Permission id
  name VARCHAR(255) UNIQUE NOT NULL, -- Permission name
  internal_name VARCHAR(255) UNIQUE NOT NULL, -- Permission internal name
  description TEXT NOT NULL NOT NULL, -- Permission description
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Creation timestamp
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL -- Update timestamp
);
CREATE INDEX idx_permissions_internal_name ON permissions(internal_name);

-- Create the 'avatars' table
CREATE TABLE IF NOT EXISTS avatars (
  id SERIAL PRIMARY KEY, -- Avatar id
  svg TEXT NOT NULL, -- Avatar base64
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Creation timestamp
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL -- Update timestamp
);
CREATE INDEX idx_avatars_svg ON avatars(svg);

-- Insert avatars
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');
INSERT INTO avatars (svg) VALUES ('<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/></svg>');

-- Insert permissions
INSERT INTO permissions (name, internal_name, description) VALUES ('Create User', 'create_user', 'Create user permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Update User', 'update_user', 'Update user permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Delete User', 'delete_user', 'Delete user permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read User', 'read_user', 'Read user permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Users', 'read_users', 'Read users permission');

INSERT INTO permissions (name, internal_name, description) VALUES ('Create Role', 'create_role', 'Create role permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Update Role', 'update_role', 'Update role permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Delete Role', 'delete_role', 'Delete role permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Role', 'read_role', 'Read role permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Roles', 'read_roles', 'Read roles permission');

INSERT INTO permissions (name, internal_name, description) VALUES ('Create Role Permission', 'create_role_permission', 'Create role permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Update Role Permission', 'update_role_permission', 'Update role permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Role Permissions', 'read_role_permissions', 'Read role permissions permission');

INSERT INTO permissions (name, internal_name, description) VALUES ('Create Permission', 'create_permission', 'Create permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Update Permission', 'update_permission', 'Update permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Delete Permission', 'delete_permission', 'Delete permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Permission', 'read_permission', 'Read permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Permissions', 'read_permissions', 'Read permissions permission');

INSERT INTO permissions (name, internal_name, description) VALUES ('Create Avatar', 'create_avatar', 'Create avatar permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Update Avatar', 'update_avatar', 'Update avatar permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Delete Avatar', 'delete_avatar', 'Delete avatar permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Avatar', 'read_avatar', 'Read avatar permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Avatars', 'read_avatars', 'Read avatars permission');

-- Create the 'users' table
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY, -- User id
  name VARCHAR(255) NOT NULL, -- User name
  email VARCHAR(255) UNIQUE NOT NULL, -- User email
  password TEXT NOT NULL, -- User password
  active BOOLEAN DEFAULT FALSE NOT NULL, -- User active
  role_id INT NOT NULL, -- User role id
  avatar_id INT NOT NULL, -- User avatar id
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Creation timestamp
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Update timestamp
  FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
  FOREIGN KEY (avatar_id) REFERENCES avatars (id) ON DELETE CASCADE
);
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_users_avatar_id ON users(avatar_id);
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_active ON users(active);

-- Insert users
INSERT INTO users (name, password, active, email, role_id, avatar_id) VALUES ('admin', '$2a$10$FcjdWT805.CjOEz9xc/P9eJojZ0.3SLlLRAgI/2ve6zPjgGY2jFsS', true, 'admin@admin.com',(SELECT id FROM roles WHERE internal_name = 'admin'), 1);
INSERT INTO users (name, password, active, email, role_id, avatar_id) VALUES ('user', '$2a$10$FcjdWT805.CjOEz9xc/P9eJojZ0.3SLlLRAgI/2ve6zPjgGY2jFsS', true, 'user@user.com',(SELECT id FROM roles WHERE internal_name = 'user'), 2);

-- Create the 'role_permission' table
CREATE TABLE IF NOT EXISTS role_permissions (
  id SERIAL PRIMARY KEY, -- Role permission id
  role_id INT NOT NULL, -- Role id
  permission_id INT NOT NULL, -- Permission id
  FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
  FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE
);
CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);

-- Insert role_permission mappings
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'create_user'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'update_user'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'delete_user'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_user'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_users'));

INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'create_role'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'update_role'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'delete_role'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_role'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_roles'));

INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'create_permission'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'update_permission'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'delete_permission'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_permission'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_permissions'));

INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'create_role_permission'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'update_role_permission'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_role_permissions'));

INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'create_avatar'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'update_avatar'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'delete_avatar'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_avatar'));
INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'admin'), (SELECT id FROM permissions WHERE internal_name = 'read_avatars'));

INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'user'), (SELECT id FROM permissions WHERE internal_name = 'read_user'));

-- Create the 'users_validation' table
CREATE TABLE IF NOT EXISTS users_validation (
  id SERIAL PRIMARY KEY, -- Validation user id
  user_id INT NOT NULL, -- User id
  hash TEXT NOT NULL, -- Validation hash
  expires_in INT NOT NULL, -- Expiration
  used BOOLEAN DEFAULT FALSE NOT NULL, -- Used
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Creation timestamp
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Update timestamp
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE INDEX idx_users_validation_user_id ON users_validation(user_id);
CREATE INDEX idx_users_validation_hash ON users_validation(hash);
CREATE INDEX idx_users_validation_used ON users_validation(used);
CREATE INDEX idx_users_validation_created_at ON users_validation(created_at);

-- Create the 'auth' table
CREATE TABLE IF NOT EXISTS auth (
  id SERIAL PRIMARY KEY, -- Auth token id
  user_id INT NOT NULL, -- User id
  token TEXT NOT NULL, -- Auth token
  refresh_token TEXT NOT NULL, -- Refresh token
  active BOOLEAN DEFAULT TRUE NOT NULL, -- Auth token active
  token_expires_in INT NOT NULL, -- Expiration
  refresh_token_expires_in INT NOT NULL, -- Expiration
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Creation timestamp
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Update timestamp
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_auth_user_id ON auth(user_id);
CREATE INDEX idx_auth_token ON auth(token);
CREATE INDEX idx_auth_refresh_token ON auth(refresh_token);
CREATE INDEX idx_auth_active ON auth(active);


-- Add table and column comments
COMMENT ON TABLE users IS 'Users table';
COMMENT ON COLUMN users.id IS 'User id';
COMMENT ON COLUMN users.name IS 'User name';
COMMENT ON COLUMN users.email IS 'User email';
COMMENT ON COLUMN users.password IS 'User password';
COMMENT ON COLUMN users.role_id IS 'User role id';

COMMENT ON TABLE auth IS 'Auth tokens table';
COMMENT ON COLUMN auth.id IS 'Auth token id';
COMMENT ON COLUMN auth.user_id IS 'User id';
COMMENT ON COLUMN auth.token IS 'Auth token';
COMMENT ON COLUMN auth.refresh_token IS 'Refresh token';
COMMENT ON COLUMN auth.active IS 'Auth token active';
COMMENT ON COLUMN auth.token_expires_in IS 'Expiration token';
COMMENT ON COLUMN auth.refresh_token_expires_in IS 'Expiration refresh token';

COMMENT ON TABLE roles IS 'Roles table';
COMMENT ON COLUMN roles.id IS 'Role id';
COMMENT ON COLUMN roles.name IS 'Role name';
COMMENT ON COLUMN roles.internal_name IS 'Role internal name';
COMMENT ON COLUMN roles.description IS 'Role description';

COMMENT ON TABLE permissions IS 'Permissions table';
COMMENT ON COLUMN permissions.id IS 'Permission id';
COMMENT ON COLUMN permissions.name IS 'Permission name';
COMMENT ON COLUMN permissions.internal_name IS 'Permission internal name';
COMMENT ON COLUMN permissions.description IS 'Permission description';

COMMENT ON TABLE role_permissions IS 'Roles and permissions table';
COMMENT ON COLUMN role_permissions.id IS 'Role permission id';
COMMENT ON COLUMN role_permissions.role_id IS 'Role id';
COMMENT ON COLUMN role_permissions.permission_id IS 'Permission id';

COMMENT ON TABLE users_validation IS 'Validation user table';
COMMENT ON COLUMN users_validation.id IS 'Validation user id';
COMMENT ON COLUMN users_validation.user_id IS 'User id';
COMMENT ON COLUMN users_validation.hash IS 'Validation hash';
COMMENT ON COLUMN users_validation.used IS 'Used';
COMMENT ON COLUMN users_validation.expires_in IS 'Expiration';

COMMENT ON TABLE avatars IS 'Avatars table';
COMMENT ON COLUMN avatars.id IS 'Avatar id';
COMMENT ON COLUMN avatars.svg IS 'Avatar base64';