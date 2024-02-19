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

INSERT INTO permissions (name, internal_name, description) VALUES ('Create Permission', 'create_permission', 'Create permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Update Permission', 'update_permission', 'Update permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Delete Permission', 'delete_permission', 'Delete permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Permission', 'read_permission', 'Read permission permission');
INSERT INTO permissions (name, internal_name, description) VALUES ('Read Permissions', 'read_permissions', 'Read permissions permission');

-- Create the 'users' table
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY, -- User id
  name VARCHAR(255) NOT NULL, -- User name
  email VARCHAR(255) UNIQUE NOT NULL, -- User email
  password TEXT NOT NULL, -- User password
  role_id INT NOT NULL, -- User role id
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Creation timestamp
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Update timestamp
  FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE
);
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_email ON users(email);

-- Insert users
INSERT INTO users (name, password, email, role_id) VALUES ('admin', '$2a$10$FcjdWT805.CjOEz9xc/P9eJojZ0.3SLlLRAgI/2ve6zPjgGY2jFsS', 'admin@admin.com',(SELECT id FROM roles WHERE internal_name = 'admin'));
INSERT INTO users (name, password, email, role_id) VALUES ('user', '$2a$10$FcjdWT805.CjOEz9xc/P9eJojZ0.3SLlLRAgI/2ve6zPjgGY2jFsS', 'user@user.com',(SELECT id FROM roles WHERE internal_name = 'user'));

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

INSERT INTO role_permissions (role_id, permission_id) VALUES ((SELECT id FROM roles WHERE internal_name = 'user'), (SELECT id FROM permissions WHERE internal_name = 'read_user'));

-- Add table and column comments
COMMENT ON TABLE users IS 'Users table';
COMMENT ON COLUMN users.id IS 'User id';
COMMENT ON COLUMN users.name IS 'User name';
COMMENT ON COLUMN users.email IS 'User email';
COMMENT ON COLUMN users.password IS 'User password';
COMMENT ON COLUMN users.role_id IS 'User role id';

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