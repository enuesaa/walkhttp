import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
  overwrite: true,
  schema: '../pkg/graph/schema.graphqls',
  generates: {
    './src/lib/graphtypes.ts': {
      plugins: ['typescript'],
    },
  },
  hooks: { afterAllFileWrite: ['prettier --write'] },
} satisfies CodegenConfig
