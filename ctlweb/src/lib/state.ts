import { atom, useAtomValue, useSetAtom } from 'jotai'
import { Invocation } from './graphtypes'

const invocationAtom = atom<undefined|Invocation>(undefined)
export const useGetInvocation = () => useAtomValue(invocationAtom)
export const useSetInvocation = () => useSetAtom(invocationAtom)
