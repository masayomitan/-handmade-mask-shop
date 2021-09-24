import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

const ItemCreate = () =>  {

  const [data, setData] = useState([])
  const [image, setImage] = useState("")

  const onSubmit = (data) => {

    const itemData = new FormData()
    itemData.append("name", data.name)
    itemData.append("detail", data.detail)
    itemData.append("image", fileInput.current.files[0])

    axios.post( '/admin/item/store', itemData, {
            headers: {
              'X-Requested-With': 'XMLHttpRequest',
              'content-type': 'multipart/form-data',
            },
        }
    )
    .then(response => {
        console.log(response.data);
    })
    .catch(err => {
        console.log(err);
    });
  };

  const handleChange = (e) => {
      const target = e.target;
      const name = target.name;
      const value = target.type === "file" ? target.files[0] : target.value;
      setData({...data, [name]: value})
  }

  const fileInput = React.createRef()

  return (
      <>
          <input 
            type="text"
            name="name"
            // value={data.name}
            required="required"
            onChange={handleChange}
          />
          <input 
            type="textarea"
            name="detail"
            // value={data.detail}
            required="required"
            onChange={handleChange}
          />
          <input 
            type="file" 
            multiple="multiple"
            name="image" 
            ref={fileInput}
            onChange={handleChange}
            accept="image/*"
          />

          <button type="submit" onClick={() => onSubmit(data)} />登録する
      </>
  )
}

const ItemForm = document.getElementById("item_create");

ReactDOM.render(
  <ItemCreate/>, ItemForm
);

// export default ItemEdit