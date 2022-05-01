import React, { useState, useEffect, useRef, useMemo, useCallback } from "react";
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'
import { useRouter } from "next/router";


const ItemsFromCategoryId = () => {

  const apiBaseUrl = process.env.NEXT_PUBLIC_BASE_URL
  const [items, setItems] = useState([])
  const [category, setCategory] = useState('')
  const [useErrorHandler, setErrorHandler] = useState(null)
  const router = useRouter();
  const id  = router.query.id;

  axios.defaults.baseURL = apiBaseUrl
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  
  const getItemsByCategoryId = useCallback(async() => {
    try {
      const result = await axios.get("/front/api/get-display-items-category/" + id)
      setItems(result.data)
    } catch(e) {
      if (axios.isAxiosError(e) && e.response && e.response.status === 400) {
        console.log('400 Error')
        setErrorHandler(() => errorHandler(res, e))
      }
    }
  }, [id])


  useEffect((id)  => {
    getItemsByCategoryId(id)
  }, [getItemsByCategoryId])


    return (
    <>
      <div>
        <ul>
        {items.map((v, i) =>
          <li key={i} className="inline">
            <a href={ "/items/detail/" + v.ID}>
              {v.ItemImages.length > 0 ?
                <Image
                  alt="image"
                  src={apiBaseUrl + v.ItemImages[0].file_path}
                  width={300}
                  height={300}
                />
                :   
                <Image
                  alt="image"
                  src={apiBaseUrl + "/public/img/no_image.png"}
                  width={300}
                  height={300}
                />
              } 
              {v.Name}
            </a>
          </li>
        )}
        </ul>
      </div>
    </>
  )
}

export default ItemsFromCategoryId