import { gql } from "@apollo/client";

export const GET_ALL_PATIENTS = gql`
  query GetAllPatients {
    patients {
      id
      name
      scans {
        id
        timeCreated
        presignedUrl
      }
    }
  }
`