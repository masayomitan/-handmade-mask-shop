import React, { useState, useEffect, useRef } from "react";
import axios from 'axios';
import Link from 'next/link';
import Header from './components/header.js';

const Top = () => {

  const [categories, setCategories] = useState([])
  
  axios.defaults.baseURL = process.env.NEXT_PUBLIC_API_BASE_URL
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
    const result = await axios.get("front/api/get-display-items");

  }
  
  useEffect(()  => {
    getCategories();
  }, [])

  return (
    <>
      <div className="">
        <Header />
      </div>
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
    
    </>
  )
}

export default Top
