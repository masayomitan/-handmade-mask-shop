import React, { useState, useEffect, useRef } from "react";
import axios from 'axios';
import Image from 'next/image';
import Link from 'next/link';
import Header from './components/header.js';
import Footer from './components/footer.js';
import ItemIndex from './components/item_index.js';

const Top = () => {
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL

  const [categories, setCategories] = useState([])

  axios.defaults.baseURL = apiBaseUrl
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  const getCategories = async () => {
    try {
      const result = await axios.get("front/api/get-categories");
      if (result.status === 200) {
        setCategories(result.data)
      }
    } catch (e) {
      if (e.response && e.response.status === 400) {
        console.log('400 Error!!')
      }
    }
  }
  
  useEffect(()  => {
    getCategories();
  }, [])

  return (
    <>
      <div className="font-serif">
        <div className="">
          <Header />
        </div>
        <div className="flex">
            <img
              className="w-full"
              alt="image"
              src={apiBaseUrl + "/public/img/tops/top.jpg"}
            />
        </div>
        <div className="flex">
          ハンドメイドを紹介ぺージ
          <h2>趣味で作ったマスクが地元で好評なのからスタート</h2>
          <a href0="/">
            楽しく趣味で作っています
          </a>
        </div>

        <div className="flex w-11/12 m-auto">
          <ul>
            {categories.map((v, i) =>
              <li key={i}>
                <Link href={{
                  pathname: "/categories/[category_id]",
                  query: {category_id: v.ID}
                }}>
                  <a>{v.name}</a>
                </Link>
              </li>
            )}
          </ul>

          <ItemIndex />
        </div>
        <div className="">
          <Footer />
        </div>
      </div>
    </>
  )
}

export default Top
