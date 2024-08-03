import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
  overwrite: true,
  schema: '../gql/*.graphql',
  documents: './src/gql/queries/*.ts',
  generates: {
    './src/gql/types.ts': {
      plugins: ['typescript'],
    },
  },
  hooks: {
    afterAllFileWrite: ['prettier --write'],
  },
} satisfies CodegenConfig
