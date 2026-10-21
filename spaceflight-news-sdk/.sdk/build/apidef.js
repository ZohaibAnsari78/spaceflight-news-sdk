
const { ApiDef } = require('@voxgig/apidef')

const opts = {
  folder: __dirname + '/../model',
  server: 'https://api.spaceflightnewsapi.net',
}

module.exports = ApiDef.makeBuild(opts)
