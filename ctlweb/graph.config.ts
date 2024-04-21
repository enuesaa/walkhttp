import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
  overwrite: true,
  schema: '../pkg/graph/schema.graphqls',
  generates: {
    './src/graph/types.ts': {
      plugins: ['typescript'],
    },
  },
  hooks: { afterAllFileWrite: ['prettier --write'] },
} satisfies CodegenConfig
