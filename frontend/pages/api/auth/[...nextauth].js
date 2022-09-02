import axios from "axios"
import React, { useState, useEffect, useRef } from "react";
import NextAuth from "next-auth"
import GoogleProvider from "next-auth/providers/google"

export default NextAuth({
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET,
    }),
    // ...add more providers here
  ],

  callbacks: {
    async session({ session, token, user }) {
      session.token = token

      //emailで検索するAPIを叩く
      //あればsessionにID格納、なければ新規登録してID格納
      const resultSession = setData(session)
      return resultSession;
    },
    async jwt({ token, user, account, profile, isNewUser }) {
      if (user) {
        token.uid = user.id;
      }
      return token;
    },
    async signIn({ user, account, profile, email, credentials }) {
      const isAllowedToSignIn = true
      console.log()

      if (isAllowedToSignIn) {
        return true
      } else {
        // Return false to display a default error message
        return false
        // Or you can return a URL to redirect to:
        // return '/unauthorized'
      }
    }
  }
})

const setData = async (session) => {
  const apiBaseUrl = process.env.NEXT_PUBLIC_BASE_URL
  axios.defaults.baseURL = apiBaseUrl
  axios.defaults.headers.common = {
    'Content-Type': 'application/sjson',
    'X-Requested-With': 'XMLHttpRequest',
  }
  const request = {
    first_name: session.user.name,
    email: session.user.email
  }

  try {
    const response = await axios.post("front/api/costomer/get-by-email", request)
    let id = response.data;
    console.log(response.data)
    session.user.id = id
  } catch (e) {
    if (e.response && e.response.status === 400) {
      console.log(e)
    }
  }
  return session
}
