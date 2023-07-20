package handler

import (
	"context"
	"net/http"

	interfaces "github.com/athunlal/bookNow-Account-Services/pkg/UseCase/interface"
	"github.com/athunlal/bookNow-Account-Services/pkg/domain"
	"github.com/athunlal/bookNow-Account-Services/pkg/pb"
)

type UserHandler struct {
	useCase interfaces.UserUseCase
	pb.ProfileManagementServer
}

func NewUserHandler(usecase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: usecase,
	}
}

func (h *UserHandler) ViewProfile(ctx context.Context, req *pb.ViewProfileRequest) (*pb.ViewProfileResponse, error) {
	user := domain.User{
		Id: uint(req.Id),
	}
	user, err := h.useCase.ViewProfile(user)
	if err != nil {
		return &pb.ViewProfileResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}
	return &pb.ViewProfileResponse{
		Status:   http.StatusOK,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Profile:  user.Profile,
		Error:    "nil",
		Dob:      user.Dateofbirth,
		Gender:   user.Gender,
	}, nil
}

func (h *UserHandler) EditProfile(ctx context.Context, req *pb.EditProfileRequest) (*pb.EditProfileResponse, error) {
	user := domain.User{
		Id:          uint(req.Id),
		Username:    req.Username,
		Dateofbirth: req.Dob,
		Gender:      req.Gender,
	}
	err := h.useCase.EditProfile(user)
	if err != nil {
		return &pb.EditProfileResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Editing the user profile",
		}, err
	} else {
		return &pb.EditProfileResponse{
			Status: http.StatusOK,
			Error:  "nil",
		}, nil
	}
}

func (h *UserHandler) ChangePassword(ctx context.Context, req *pb.ChangeRequest) (*pb.ChangeResponse, error) {
	passwordData := domain.Password{
		Id:          uint(req.Id),
		Oldpassword: req.Oldpassword,
		Newpassword: req.Newpassword,
	}
	err := h.useCase.ChangePassword(passwordData)
	if err != nil {
		return &pb.ChangeResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Changing the Password",
		}, err
	} else {
		return &pb.ChangeResponse{
			Status: http.StatusOK,
			Error:  "nil",
		}, nil
	}
}

// ------------------------------------ Address Management ------------------------------------------

func (h *UserHandler) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	addressData := domain.Address{
		Uid:             uint(req.Id),
		Type:            req.Type,
		Locationaddress: req.Locationaddress,
		CompleteAddress: req.Completeaddress,
		Landmark:        req.Landmark,
		Floorno:         req.Floorno,
	}
	addressData, err := h.useCase.AddAddress(addressData)
	if err != nil {
		return &pb.AddAddressResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Adding the Address",
		}, err
	} else {
		return &pb.AddAddressResponse{
			Status: http.StatusOK,
			Error:  "nil",
			Addid:  int64(addressData.Addressid),
		}, nil
	}

}

func (h *UserHandler) ViewAddress(ctx context.Context, req *pb.ViewAddressRequest) (*pb.ViewAddressResponse, error) {
	addressData := domain.Address{
		Uid: uint(req.Id),
	}
	var address []domain.Address
	address, err := h.useCase.ViewAddress(addressData)
	if err != nil {
		return &pb.ViewAddressResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Viewing the address",
		}, err
	}

	var pbAddresses []*pb.Address
	for _, addr := range address {
		pbAddr := &pb.Address{
			Addressid:       int64(addr.Addressid),
			Type:            addr.Type,
			Locationaddress: addr.Locationaddress,
			Completeaddress: addr.CompleteAddress,
			Landmark:        addr.Landmark,
			Floorno:         addr.Floorno,
		}
		pbAddresses = append(pbAddresses, pbAddr)
	}
	return &pb.ViewAddressResponse{
		Status:    http.StatusOK,
		Addresses: pbAddresses,
		Error:     "nil",
	}, nil
}

func (h *UserHandler) ViewAddressById(ctx context.Context, req *pb.ViewAddressByIdRequest) (*pb.ViewAddressByIdResponse, error) {
	addressData := domain.Address{
		Addressid: uint(req.Addid),
		Uid:       uint(req.Uid),
	}
	addressData, err := h.useCase.ViewAddressByID(addressData)
	if err != nil {
		return &pb.ViewAddressByIdResponse{
			Status: http.StatusNotFound,
			Error:  "Error in Viewing the address",
		}, err
	} else {
		return &pb.ViewAddressByIdResponse{
			Status:          http.StatusOK,
			Error:           "nil",
			Addressid:       int64(addressData.Addressid),
			Type:            addressData.Type,
			Locationaddress: addressData.Locationaddress,
			Completeaddress: addressData.CompleteAddress,
			Landmark:        addressData.Landmark,
			Floorno:         addressData.Floorno,
		}, err
	}

}

func (h *UserHandler) EditAddress(ctx context.Context, req *pb.EditAddressRequest) (*pb.EditAddressResponse, error) {
	addreddData := domain.Address{
		Addressid:       uint(req.Addressid),
		Uid:             uint(req.Id),
		Type:            req.Type,
		Locationaddress: req.Locationaddress,
		CompleteAddress: req.Completeaddress,
		Landmark:        req.Landmark,
		Floorno:         req.Floorno,
	}

	addreddData, err := h.useCase.EditAddress(addreddData)
	if err != nil {
		return &pb.EditAddressResponse{
			Status: http.StatusBadRequest,
			Error:  "Error in Editing the Address",
			Addid:  int64(addreddData.Addressid),
		}, err
	} else {
		return &pb.EditAddressResponse{
			Status: http.StatusOK,
			Error:  "nil",
			Addid:  int64(addreddData.Addressid),
		}, nil
	}
}
