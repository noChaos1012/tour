package errcode

import (
	pb "github.com/noChaos1012/tour/tag-service/proto"
	code "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(err *Error) error {
	s, _ := status.New((TogRPCCode(err.code)), err.Msg()).WithDetails(&pb.Error{Code: int32(err.Code()), Message: err.Msg()})
	return s.Err()
}

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func TogRPCStatus(c int, msg string) *Status {
	s, _ := status.New(TogRPCCode(c), msg).WithDetails(&pb.Error{Code: int32(c), Message: msg})
	return &Status{s}
}

func TogRPCCode(c int) code.Code {
	switch c {
	case Success.code:
		return code.OK

	case Fail.code:
		return code.Internal

	case InvalidParams.code:
		return code.InvalidArgument

	case Unauthorized.code:
		return code.Unauthenticated

	case AccessDenied.code:
		return code.PermissionDenied

	case DeadlineExceeded.code:
		return code.DeadlineExceeded

	case NotFound.code:
		return code.NotFound

	case LimitExceed.code:
		return code.ResourceExhausted
	case MethodNotAllowed.code:
		return code.Unimplemented
	}

	return code.Unknown
}
