export const ShowTagSpan = ({tags}: {tags: string[]}) => {
  console.log(tags)
  return (
    <>
      {tags.length > 1 &&
        tags.map(tag => <span className='border border-success rounded mx-1 p-1'>{tag}</span>)
      }
    </>
  )
}