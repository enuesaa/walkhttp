type Props = {
  title: string
}
export const MakeFormHeading = ({ title }: Props) => {
  return (
    <h4 className='block text-lg pt-2 px-3 font-medium text-stone-300'>
      {title}
    </h4>
  )
}
