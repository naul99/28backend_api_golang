// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/movie.proto

package movie

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Movie struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OriginalName  string                 `protobuf:"bytes,3,opt,name=originalName,proto3" json:"originalName,omitempty"`
	Slug          string                 `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Age           string                 `protobuf:"bytes,5,opt,name=age,proto3" json:"age,omitempty"`
	Runtime       string                 `protobuf:"bytes,6,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Quality       string                 `protobuf:"bytes,7,opt,name=quality,proto3" json:"quality,omitempty"`
	Rating        string                 `protobuf:"bytes,8,opt,name=rating,proto3" json:"rating,omitempty"`
	Imdb          string                 `protobuf:"bytes,9,opt,name=imdb,proto3" json:"imdb,omitempty"`
	Tmdb          string                 `protobuf:"bytes,10,opt,name=tmdb,proto3" json:"tmdb,omitempty"`
	RelaseDate    string                 `protobuf:"bytes,11,opt,name=relaseDate,proto3" json:"relaseDate,omitempty"`
	Trailer       string                 `protobuf:"bytes,12,opt,name=trailer,proto3" json:"trailer,omitempty"`
	EpisdoTotal   string                 `protobuf:"bytes,13,opt,name=episdoTotal,proto3" json:"episdoTotal,omitempty"`
	Season        string                 `protobuf:"bytes,14,opt,name=season,proto3" json:"season,omitempty"`
	Images        []string               `protobuf:"bytes,15,rep,name=images,proto3" json:"images,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Movie) Reset() {
	*x = Movie{}
	mi := &file_api_movie_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *Movie) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Movie) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *Movie) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Movie) GetQuality() string {
	if x != nil {
		return x.Quality
	}
	return ""
}

func (x *Movie) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *Movie) GetImdb() string {
	if x != nil {
		return x.Imdb
	}
	return ""
}

func (x *Movie) GetTmdb() string {
	if x != nil {
		return x.Tmdb
	}
	return ""
}

func (x *Movie) GetRelaseDate() string {
	if x != nil {
		return x.RelaseDate
	}
	return ""
}

func (x *Movie) GetTrailer() string {
	if x != nil {
		return x.Trailer
	}
	return ""
}

func (x *Movie) GetEpisdoTotal() string {
	if x != nil {
		return x.EpisdoTotal
	}
	return ""
}

func (x *Movie) GetSeason() string {
	if x != nil {
		return x.Season
	}
	return ""
}

func (x *Movie) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type Genre struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slug          string                 `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Genre) Reset() {
	*x = Genre{}
	mi := &file_api_movie_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{1}
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Genre) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type Episode struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServerId      string                 `protobuf:"bytes,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Episode       string                 `protobuf:"bytes,2,opt,name=episode,proto3" json:"episode,omitempty"`
	Hls           string                 `protobuf:"bytes,3,opt,name=hls,proto3" json:"hls,omitempty"`
	Servers       []*Server              `protobuf:"bytes,4,rep,name=servers,proto3" json:"servers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Episode) Reset() {
	*x = Episode{}
	mi := &file_api_movie_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Episode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Episode) ProtoMessage() {}

func (x *Episode) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Episode.ProtoReflect.Descriptor instead.
func (*Episode) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{2}
}

func (x *Episode) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Episode) GetEpisode() string {
	if x != nil {
		return x.Episode
	}
	return ""
}

func (x *Episode) GetHls() string {
	if x != nil {
		return x.Hls
	}
	return ""
}

func (x *Episode) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_api_movie_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{3}
}

func (x *Server) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMovieListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Keyword       string                 `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMovieListRequest) Reset() {
	*x = GetMovieListRequest{}
	mi := &file_api_movie_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMovieListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieListRequest) ProtoMessage() {}

func (x *GetMovieListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieListRequest.ProtoReflect.Descriptor instead.
func (*GetMovieListRequest) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{4}
}

func (x *GetMovieListRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *GetMovieListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetMovieListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMovieListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMovieListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Movies        []*Movie               `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMovieListResponse) Reset() {
	*x = GetMovieListResponse{}
	mi := &file_api_movie_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMovieListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieListResponse) ProtoMessage() {}

func (x *GetMovieListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieListResponse.ProtoReflect.Descriptor instead.
func (*GetMovieListResponse) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{5}
}

func (x *GetMovieListResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

type GetMovieDetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Slug          string                 `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMovieDetailRequest) Reset() {
	*x = GetMovieDetailRequest{}
	mi := &file_api_movie_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMovieDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieDetailRequest) ProtoMessage() {}

func (x *GetMovieDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieDetailRequest.ProtoReflect.Descriptor instead.
func (*GetMovieDetailRequest) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{6}
}

func (x *GetMovieDetailRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetMovieDetailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Movie         *Movie                 `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	Genres        []*Genre               `protobuf:"bytes,2,rep,name=genres,proto3" json:"genres,omitempty"`
	Episodes      []*Episode             `protobuf:"bytes,3,rep,name=episodes,proto3" json:"episodes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMovieDetailResponse) Reset() {
	*x = GetMovieDetailResponse{}
	mi := &file_api_movie_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMovieDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieDetailResponse) ProtoMessage() {}

func (x *GetMovieDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_movie_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMovieDetailResponse) Descriptor() ([]byte, []int) {
	return file_api_movie_proto_rawDescGZIP(), []int{7}
}

func (x *GetMovieDetailResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

func (x *GetMovieDetailResponse) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *GetMovieDetailResponse) GetEpisodes() []*Episode {
	if x != nil {
		return x.Episodes
	}
	return nil
}

var File_api_movie_proto protoreflect.FileDescriptor

var file_api_movie_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70, 0x62, 0x22, 0xf5, 0x02, 0x0a, 0x05, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6d, 0x64, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d,
	0x64, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6d, 0x64, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x6d, 0x64, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x65, 0x70, 0x69, 0x73, 0x64, 0x6f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x70, 0x69, 0x73, 0x64, 0x6f, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x2f, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x22, 0x7c, 0x0a, 0x07, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x68, 0x6c, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x22, 0x2c, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x71, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x3e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22,
	0x94, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x65, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x32, 0xae, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_movie_proto_rawDescOnce sync.Once
	file_api_movie_proto_rawDescData []byte
)

func file_api_movie_proto_rawDescGZIP() []byte {
	file_api_movie_proto_rawDescOnce.Do(func() {
		file_api_movie_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_movie_proto_rawDesc), len(file_api_movie_proto_rawDesc)))
	})
	return file_api_movie_proto_rawDescData
}

var file_api_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_movie_proto_goTypes = []any{
	(*Movie)(nil),                  // 0: moviepb.Movie
	(*Genre)(nil),                  // 1: moviepb.Genre
	(*Episode)(nil),                // 2: moviepb.Episode
	(*Server)(nil),                 // 3: moviepb.Server
	(*GetMovieListRequest)(nil),    // 4: moviepb.GetMovieListRequest
	(*GetMovieListResponse)(nil),   // 5: moviepb.GetMovieListResponse
	(*GetMovieDetailRequest)(nil),  // 6: moviepb.GetMovieDetailRequest
	(*GetMovieDetailResponse)(nil), // 7: moviepb.GetMovieDetailResponse
}
var file_api_movie_proto_depIdxs = []int32{
	3, // 0: moviepb.Episode.servers:type_name -> moviepb.Server
	0, // 1: moviepb.GetMovieListResponse.movies:type_name -> moviepb.Movie
	0, // 2: moviepb.GetMovieDetailResponse.movie:type_name -> moviepb.Movie
	1, // 3: moviepb.GetMovieDetailResponse.genres:type_name -> moviepb.Genre
	2, // 4: moviepb.GetMovieDetailResponse.episodes:type_name -> moviepb.Episode
	4, // 5: moviepb.MovieService.GetMovieList:input_type -> moviepb.GetMovieListRequest
	6, // 6: moviepb.MovieService.GetMovieDetail:input_type -> moviepb.GetMovieDetailRequest
	5, // 7: moviepb.MovieService.GetMovieList:output_type -> moviepb.GetMovieListResponse
	7, // 8: moviepb.MovieService.GetMovieDetail:output_type -> moviepb.GetMovieDetailResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_movie_proto_init() }
func file_api_movie_proto_init() {
	if File_api_movie_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_movie_proto_rawDesc), len(file_api_movie_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_movie_proto_goTypes,
		DependencyIndexes: file_api_movie_proto_depIdxs,
		MessageInfos:      file_api_movie_proto_msgTypes,
	}.Build()
	File_api_movie_proto = out.File
	file_api_movie_proto_goTypes = nil
	file_api_movie_proto_depIdxs = nil
}
