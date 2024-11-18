type Props = {
  title: string
}
export const HistoryDialogHeading = ({ title }: Props) => {
  return (
    <h4 className='block text-lg text-center py-1 font-medium text-stone-300'>
      {title}
    </h4>
  )
}
