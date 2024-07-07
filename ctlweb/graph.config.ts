import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
  overwrite: true,
  schema: '../pkg/invoke/schema/schema.graphqls',
  generates: {
    './src/graph/types.ts': {
      plugins: ['typescript'],
    },
  },
  hooks: { afterAllFileWrite: ['prettier --write'] },
} satisfies CodegenConfig
