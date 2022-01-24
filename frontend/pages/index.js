import React, { useState, useEffect, useRef } from "react";
import axios from 'axios';
import Image from 'next/image';
import Link from 'next/link';
import Header from './components/header.js';
import ItemIndex from './components/item_index.js';

const Top = () => {
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL

  const [items, setItems] = useState([])
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

  const getItems = async () => {
    try {
    const result = await axios.get("front/api/get-display-items");
    setItems(result.data)
    } catch (e) {
      if (axios.isAxiosError(e) && e.response && e.response.status === 400) {
        console.log('400 Error!!')
      }
      console.log(e)
    }
  }
  
  useEffect(()  => {
    getCategories();
    getItems();
  }, [])

  return (
    <>
      <div className="">
        <Header />
      </div>
      <div className="flex">
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
    </>
  )
}

export default Top
