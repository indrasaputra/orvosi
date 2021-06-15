package entity

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	orvosiv1 "github.com/indrasaputra/orvosi/proto/indrasaputra/orvosi/v1"
)

// ErrInternal returns codes.Internal explained that unexpected behavior occurred in system.
func ErrInternal(message string) error {
	st := status.New(codes.Internal, message)
	te := &orvosiv1.MedicalRecordError{
		ErrorCode: orvosiv1.MedicalRecordErrorCode_INTERNAL,
	}
	res, err := st.WithDetails(te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrEmptyMedicalRecord returns codes.InvalidArgument explained that the instance is empty or nil.
func ErrEmptyMedicalRecord() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "Medical record instance",
		Description: "empty or nil",
	})

	te := &orvosiv1.MedicalRecordError{
		ErrorCode: orvosiv1.MedicalRecordErrorCode_EMPTY_MEDICAL_RECORD,
	}
	res, err := st.WithDetails(br, te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

func createBadRequest(details ...*errdetails.BadRequest_FieldViolation) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: details,
	}
}
