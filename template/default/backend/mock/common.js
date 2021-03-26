const Mock = require("mockjs")

const tokens = {
  admin: {
    token: 'admin-token'
  },
  editor: {
    token: 'editor-token'
  }
}

const users = {
  'admin-token': {
    roles: ['admin'],
    introduction: 'I am a super administrator',
    avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
    name: 'Super Admin'
  },
  'editor-token': {
    roles: ['editor'],
    introduction: 'I am an editor',
    avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
    name: 'Normal Editor'
  }
}

module.exports = [
  // user login
  {
    url: '/user/login',
    type: 'post',
    response: config => {
      const {username} = config.body
      const token = tokens[username]

      // mock error
      if (!token) {
        return {
          code: 60204,
          msg: '账号或者密码不正确'
        }
      }

      return {
        code: 20000,
        data: token
      }
    }
  },

  // get user info
  {
    url: '/user/info\.*',
    type: 'get',
    response: config => {
      const {token} = config.query
      const info = users[token]

      // mock error
      if (!info) {
        return {
          code: 50008,
          msg: '登录失败，未能获取用户详情'
        }
      }

      return {
        code: 20000,
        data: info
      }
    }
  },

  // user logout
  {
    url: '/user/logout',
    type: 'post',
    response: _ => {
      return {
        code: 20000,
        data: 'success'
      }
    }
  },

  // find dict
  {
    url: '/dict/find-dict-item-list',
    type: 'post',
    response: _ => {
      return {
        code: 20000,
        data: Mock.mock({
          total: 10,
          'items|5-10': [{
            'ID|+1': 1,
            name: "@string(4,16)",
            title: "@ctitle",
            'value|+10': 10,
            description: "@string(1,32)",
            isDefault: "@boolean",
            'tagType|1': ["success", "info", "warning", "danger"],
            'status|1': [10, 20],
          }]
        })
      }
    }
  },
  // file upload
  {
    url: '/file/upload',
    type: 'post',
    response: _ => {
      return {
        code: 20000,
        data: Mock.mock({
          total: 10,
          // 属性 list 的值是一个数组，其中含有 1 到 10 个元素
          'items|1-10': [{
            'ID|+1': 1,
            filename: "@string(1,10)",
            URL: "@image"
          }]
        })
      }
    }
  }
]
