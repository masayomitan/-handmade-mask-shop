import Link from 'next/link'

const Footer = () => {
  return (
    <>
      <div className="flex justify-around mt-4 mb-4 ">

        <ul className="flex justify-around items-center">
          <li className="mr-4">
            <Link href="/others/about">
              <a>このサイトについて</a>
            </Link>
          </li>
          <li>
            <Link href="/others/info">
              <a>お知らせ</a>
            </Link>
          </li>
        </ul>
      </div>
    </>
  )
}

export default Footer