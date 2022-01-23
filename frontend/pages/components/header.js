import Link from 'next/link'

const Header = () => {
  return (
    <>
      <div className="flex justify-around mt-4 mb-4 ">
        <h1 className="text-3xl font-bold">
          handmade shop
        </h1>

        <ul className="flex justify-around items-center">
          <li className="mr-4">
            <Link href="/others/about">
              <a>このサイトについて</a>
            </Link>
          </li>
          <li>
            <Link href="/others/notice/index">
              <a>お知らせ</a>
            </Link>
          </li>
        </ul>
      </div>
    </>
  )
}

export default Header