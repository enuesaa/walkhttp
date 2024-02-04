import { Box, Button, Select, Text, TextField } from '@radix-ui/themes'
import styles from './RequestBox.css'
import { InvokeReq, useInvoke } from '@/lib/api'
import { useForm } from 'react-hook-form'

export const RequestBox = () => {
  const invoke = useInvoke()
  const form = useForm<InvokeReq>()

  const hanldeSubmit = form.handleSubmit(async (data) => {
    await invoke.mutateAsync(data)
    form.reset()
  })

  return (
    <Box grow='1' shrink='0' className={styles.main}>
      request
      <form onSubmit={hanldeSubmit}>
        <label>
          <Text as='div' size='2' mb='1' weight='bold'>Method</Text>
          <Select.Root defaultValue='GET' {...form.register('method')}>
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
          <TextField.Input {...form.register('url')} />
        </label>
        <label>
          <Text as='div' size='2' mb='1' weight='bold'>Body</Text>
          <TextField.Input {...form.register('body')} />
        </label>
        <Button type='submit'>send</Button>
      </form>
      <Text>
        {JSON.stringify(invoke.data)}
      </Text>
    </Box>
  )
}
