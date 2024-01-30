import { useMutation, useQuery } from 'react-query'

export type Invocation = {
  id: number;
  method: string;
  url: string;
}
export const useListInvocations = () => useQuery('listInvocations', async (): Promise<Invocation[]> => {
  const res = await fetch('/api/invocations')
  const body = await res.json()
  return body?.items
})

export type InvokeReq = {
  method: string;
  url: string;
  body: string;
}
export const useInvoke = () => useMutation({
  mutationFn: async (req: InvokeReq) => {
    const res = await fetch('/api/invoke', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(req),
    })
    console.log(res)
  },
})
