import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
  overwrite: true,
  schema: '../pkg/serve/schema/*.graphql',
  documents: './src/gqlqueries/*.ts',
  generates: {
    './src/graph/types.ts': {
      plugins: ['typescript'],
    },
  },
  hooks: {
    afterAllFileWrite: ['prettier --write'],
  },
} satisfies CodegenConfig
