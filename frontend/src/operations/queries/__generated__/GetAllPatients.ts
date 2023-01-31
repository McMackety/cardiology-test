/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetAllPatients
// ====================================================

export interface GetAllPatients_patients_scans {
  __typename: "Scan";
  id: string;
  timeCreated: number;
  presignedUrl: string;
}

export interface GetAllPatients_patients {
  __typename: "Patient";
  id: string;
  name: string;
  scans: GetAllPatients_patients_scans[];
}

export interface GetAllPatients {
  patients: GetAllPatients_patients[];
}
