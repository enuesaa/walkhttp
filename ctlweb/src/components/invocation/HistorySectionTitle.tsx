import { Heading, Separator } from '@radix-ui/themes'

type Props = {
  title: string
}
export const HistorySectionTitle = ({ title }: Props) => {
  return (
    <Heading ml='3' my='2' size='3' color='gray' weight='regular'>
      {title}
    </Heading>
  )
}
