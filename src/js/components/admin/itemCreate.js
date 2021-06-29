import React from "react";
import ReactDOM from "react-dom";
import axios from 'axios';
import { useForm, SubmitHandler, SubmitErrorHandler } from 'react-hook-form'

const ItemCreate = () =>  {
  const { register, handleSubmit, watch, formState: { errors } } = useForm();

  const onSubmit = (data) => {
    console.log(data)
    
    axios.post( '/admin/item/store', data )
        .then(response => {
          console.log(response.data);
      })
      .catch(err => {
          console.log(err);
      });

  }


  return (
  
    <form onSubmit={handleSubmit(onSubmit)}>
      {/* <input defaultValue="test" {...register("name")} /> */}
      
      <input {...register("name", { required: true })} />

      {errors.name && <span>未入力です</span>}
      
      <input type="submit" />
    </form>
  )
}

const ItemForm = document.getElementById("item_create");

ReactDOM.render(

  <ItemCreate/>, ItemForm
  
);