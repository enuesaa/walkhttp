export const fmtdate = (datestr: string): string => {
  const datefmt = new Intl.DateTimeFormat('en-US', {
    hour: 'numeric',
    hour12: false,
    minute: 'numeric',
    second: 'numeric',
  })
  try {
    const date = new Date(datestr)
    return datefmt.format(date)
  } catch {
    return datestr
  }
}
