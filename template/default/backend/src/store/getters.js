const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.common.token,
  avatar: state => state.common.avatar,
  name: state => state.common.name,
  roles: state => state.common.roles,
  permission_routes: state => state.permission.routes
}
export default getters
