import Link from 'next/link'

const Header = () => {
  return (
    <>
      <div className="flex justify-around mt-4 mb-4 ">
        <h1 className="text-3xl font-bold">
          handmade shop
        </h1>
        <p>
          <Link href={"/api/auth/signin"}>
            ログイン
          </Link>
        </p>

        <div>
            <input 
              type="text" 
              name="search" 

            />
            <button type="submit" >
              検索
            </button>
        </div>
      </div>
    </>
  )
}

export default Header