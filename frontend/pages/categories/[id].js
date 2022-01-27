import React, { useState, useEffect, useRef, useMemo, useCallback } from "react";
import axios from 'axios';
import { useRouter } from "next/router";


const ItemsFromCategoryId = () => {

  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL
  
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
  
  const getDisplayItemsCategoryId = useCallback(async() => {
    try {
      const result = await axios.get("/front/api/get-display-items-category/" + id)
console.log(result.data)
    } catch(e) {
      console.log(e)
      if (axios.isAxiosError(e) && e.response && e.response.status === 400) {
        console.log('400 Error!!')
        setErrorHandler(() => errorHandler(res, e))
      }
    }
  }, [id])


  useEffect((id)  => {
    getDisplayItemsCategoryId(id)
  }, [getDisplayItemsCategoryId])


    return (
    <>

    </>
  )
}

export default ItemsFromCategoryId