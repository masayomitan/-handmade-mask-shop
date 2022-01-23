import React, { useState, useEffect, useRef } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';


const Categories = () =>  {
  const input = useRef(null)
  const [data, setData] = useState([])
  const [category, setCategory] = useState()
  const [categories, setCategories] = useState([])
  const [response, setResponse] = useState()
  // let csrfToken = document.getElementsByName("csrf_token")[0].value

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
          <span className="">1~20文字で入力してください</span>
        </>
      )
    }
    

  const postCategory = () => {
    const name = input.current.value;
    console.log(name)
    axios.post("/admin/api/post-category", {
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

  const updateCategory = (id, category) => {
    if (category.length === 0) {
      return false
    }
    if (id != category.id) {
      return false
    }

    axios.post("/admin/api/update-category/"+ id, {
      name: category.name
    })
    .then(response => {
      const res = response.data
      getCategories()
    })
    .catch(error => {
      console.log(error)
    })
  }

  const deleteCategory = async (id) => {
    console.log(id)
    try {
      const canDelete = await axios.get("/admin/api/delete-category/" + id)
      if (canDelete.data === false) {
        alert("このカテゴリーは他の商品と既に登録されています")
      }
      getCategories()
    } 
    catch (error) {
      console.log(error)
    }
  }
  
  const handleChange = (e) => {
    const t = e.target
    const diff = new Array
    diff["id"] = t.id
    diff["name"] = t.value
    setData({...data, ...diff})
    return
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
                id={v.ID}
                defaultValue={v.name}
                onChange={handleChange}
                />
                
                </label>
                <button className="" onClick={() => updateCategory(v.ID, data)}>登録</button>
                <button className="" onClick={() => deleteCategory(v.ID)}>削除</button>
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

if (document.getElementById("get_categories")) {
  ReactDOM.render(
    <Categories />, document.getElementById("get_categories")
  );
}