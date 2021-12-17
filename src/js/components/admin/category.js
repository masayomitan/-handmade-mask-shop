import React, { useState, useEffect, useRef } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';


const Categories = () =>  {
  const input = useRef(null)
  const [category, setCategory] = useState()
  const [categories, setCategories] = useState([])
  const [response, setResponse] = useState()
  // let csrfToken = document.getElementsByName("csrf_token")[0].value
  //configとして情報を設置する
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  useEffect(()  => {
    getCategories();
    
  }, [])

    const getCategories = async () => {
      try {
        const result = await axios.get("/admin/api/get-categories");
        const getCategoryData = result.data
        setCategories(getCategoryData)
        
      } catch (e) {
        if (e.response && e.response.status === 400) {
          console.log('400 Error!!')
        }
      }
    }

    const categoryVailidate = () => {
      return (
        <>
          <span className="">1~20にしてもらっていいかな？</span>
        </>
      )
    }
    

  const postCategory = () => {
    const name = input.current.value;
    console.log(name)
    axios.post("/admin/api/post-categories", {
      name: name
    })
    .then(response => {
      const res = response.data
      getCategories()
      input.current.value = ""
    })
    .catch(error => {
      console.log(error)
    })
  }

  return (
    <>
    <div className="select">
        <label>カテゴリー</label>
        <ul>
          {categories.map((v, i) =>
            <li key={i}>
              <label>
              <input
                type="text"
                defaultValue={v.name}
                />
                x
                </label>
            </li>
          )}
        </ul>
        <input
            ref={input}
            type="text"
            name="name"
            value={category}
        />
        <button className="btn" onClick={() => postCategory()}>
        作成
        </button>
        </div>
    </>
  )
}

  ReactDOM.render(
    <Categories />, document.getElementById("get_categories")
  );