import { Box, Button, Select, Text, TextField } from '@radix-ui/themes'
import styles from './RequestBox.css'
import { InvokeReq, useInvoke } from '@/lib/api'
import { useForm } from 'react-hook-form'

export const RequestBox = () => {
  const invoke = useInvoke()
  const form = useForm<InvokeReq>()

  const hanldeSubmit = form.handleSubmit(async (data) => {
    if (data.method === undefined) {
      data.method = 'GET'
    }
    await invoke.mutateAsync(data)
    form.reset()
  })

  return (
    <Box flexGrow='1' flexShrink='0' className={styles.main}>
      request
      <form onSubmit={hanldeSubmit}>
        <label>
          <Text as='div' size='2' mb='1' weight='bold'>Method</Text>
          <Select.Root defaultValue='GET'
            {...form.register('method')}
            // https://stackoverflow.com/questions/75815473/
            onValueChange={value => form.setValue('method', value)}
          >
            <Select.Trigger />
            <Select.Content>
              <Select.Item value='GET'>GET</Select.Item>
              <Select.Item value='POST'>POST</Select.Item>
              <Select.Item value='PUT'>PUT</Select.Item>
              <Select.Item value='DELETE'>DELETE</Select.Item>
            </Select.Content>
          </Select.Root>
        </label>
        <label>
          <Text as='div' size='2' mb='1' weight='bold'>Url</Text>
          <TextField.Root type='url' {...form.register('url')} />
        </label>
        <label>
          <Text as='div' size='2' mb='1' weight='bold'>Body</Text>
          <TextField.Root {...form.register('body')} />
        </label>
        <Button type='submit'>send</Button>
      </form>
      <Text>
        {JSON.stringify(invoke.data)}
      </Text>
    </Box>
  )
}
