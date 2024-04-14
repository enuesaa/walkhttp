import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
  overwrite: true,
  schema: '../pkg/graph/*.graphqls',
  generates: {
    './src/lib/graphtypes.ts': {
      plugins: ['typescript'],
    },
  },
} satisfies CodegenConfig
