package mock_dal

import (
	"context"
	"errors"
	"testing"

	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/update"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewMockReadSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadSession(ctrl)
	assert.NotNil(t, mockSession)
	assert.NotNil(t, mockSession.EXPECT())
}

func TestMockReadSession_Exists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadSession(ctrl)
	ctx := context.Background()
	key := &dal.Key{}

	t.Run("exists returns true", func(t *testing.T) {
		mockSession.EXPECT().Exists(ctx, key).Return(true, nil)

		exists, err := mockSession.Exists(ctx, key)
		assert.True(t, exists)
		assert.NoError(t, err)
	})

	t.Run("exists returns error", func(t *testing.T) {
		expectedErr := errors.New("exists error")
		mockSession.EXPECT().Exists(ctx, key).Return(false, expectedErr)

		exists, err := mockSession.Exists(ctx, key)
		assert.False(t, exists)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockReadSession_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadSession(ctrl)
	ctx := context.Background()

	t.Run("get success", func(t *testing.T) {
		mockSession.EXPECT().Get(ctx, gomock.Any()).Return(nil)

		err := mockSession.Get(ctx, &mockRecord{})
		assert.NoError(t, err)
	})

	t.Run("get error", func(t *testing.T) {
		expectedErr := errors.New("get error")
		mockSession.EXPECT().Get(ctx, gomock.Any()).Return(expectedErr)

		err := mockSession.Get(ctx, &mockRecord{})
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockReadSession_GetMulti(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadSession(ctrl)
	ctx := context.Background()
	records := []dal.Record{&mockRecord{}}

	t.Run("get multi success", func(t *testing.T) {
		mockSession.EXPECT().GetMulti(ctx, gomock.Any()).Return(nil)

		err := mockSession.GetMulti(ctx, records)
		assert.NoError(t, err)
	})

	t.Run("get multi error", func(t *testing.T) {
		expectedErr := errors.New("get multi error")
		mockSession.EXPECT().GetMulti(ctx, gomock.Any()).Return(expectedErr)

		err := mockSession.GetMulti(ctx, records)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockReadSession_QueryAllRecords(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadSession(ctrl)
	ctx := context.Background()

	t.Run("query all records success", func(t *testing.T) {
		expectedRecords := []dal.Record{&mockRecord{}}
		mockSession.EXPECT().QueryAllRecords(ctx, gomock.Any()).Return(expectedRecords, nil)

		records, err := mockSession.QueryAllRecords(ctx, nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedRecords, records)
	})

	t.Run("query all records error", func(t *testing.T) {
		expectedErr := errors.New("query error")
		mockSession.EXPECT().QueryAllRecords(ctx, gomock.Any()).Return(nil, expectedErr)

		records, err := mockSession.QueryAllRecords(ctx, nil)
		assert.Error(t, err)
		assert.Nil(t, records)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockReadSession_QueryReader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadSession(ctrl)
	ctx := context.Background()

	t.Run("query reader success", func(t *testing.T) {
		mockSession.EXPECT().QueryReader(ctx, gomock.Any()).Return(nil, nil)

		reader, err := mockSession.QueryReader(ctx, nil)
		assert.NoError(t, err)
		assert.Nil(t, reader)
	})

	t.Run("query reader error", func(t *testing.T) {
		expectedErr := errors.New("query reader error")
		mockSession.EXPECT().QueryReader(ctx, gomock.Any()).Return(nil, expectedErr)

		reader, err := mockSession.QueryReader(ctx, nil)
		assert.Error(t, err)
		assert.Nil(t, reader)
		assert.Equal(t, expectedErr, err)
	})
}

func TestNewMockWriteSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockWriteSession(ctrl)
	assert.NotNil(t, mockSession)
	assert.NotNil(t, mockSession.EXPECT())
}

func TestMockWriteSession_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockWriteSession(ctrl)
	ctx := context.Background()
	key := &dal.Key{}

	t.Run("delete success", func(t *testing.T) {
		mockSession.EXPECT().Delete(ctx, key).Return(nil)

		err := mockSession.Delete(ctx, key)
		assert.NoError(t, err)
	})

	t.Run("delete error", func(t *testing.T) {
		expectedErr := errors.New("delete error")
		mockSession.EXPECT().Delete(ctx, key).Return(expectedErr)

		err := mockSession.Delete(ctx, key)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("delete multi success", func(t *testing.T) {
		keys := []*dal.Key{key}
		mockSession.EXPECT().DeleteMulti(ctx, gomock.Any()).Return(nil)
		err := mockSession.DeleteMulti(ctx, keys)
		assert.NoError(t, err)
	})

	t.Run("delete multi error", func(t *testing.T) {
		keys := []*dal.Key{key}
		expectedErr := errors.New("delete multi error")
		mockSession.EXPECT().DeleteMulti(ctx, gomock.Any()).Return(expectedErr)
		err := mockSession.DeleteMulti(ctx, keys)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockWriteSession_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockWriteSession(ctrl)
	ctx := context.Background()

	t.Run("insert success", func(t *testing.T) {
		mockSession.EXPECT().Insert(ctx, gomock.Any()).Return(nil)

		err := mockSession.Insert(ctx, &mockRecord{})
		assert.NoError(t, err)
	})

	t.Run("insert error", func(t *testing.T) {
		expectedErr := errors.New("insert error")
		mockSession.EXPECT().Insert(ctx, gomock.Any()).Return(expectedErr)

		err := mockSession.Insert(ctx, &mockRecord{})
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("insert multi success", func(t *testing.T) {
		records := []dal.Record{&mockRecord{}}
		mockSession.EXPECT().InsertMulti(ctx, gomock.Any()).Return(nil)
		err := mockSession.InsertMulti(ctx, records)
		assert.NoError(t, err)
	})

	t.Run("insert multi error", func(t *testing.T) {
		records := []dal.Record{&mockRecord{}}
		expectedErr := errors.New("insert multi error")
		mockSession.EXPECT().InsertMulti(ctx, gomock.Any()).Return(expectedErr)
		err := mockSession.InsertMulti(ctx, records)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockWriteSession_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockWriteSession(ctrl)
	ctx := context.Background()

	t.Run("set success", func(t *testing.T) {
		mockSession.EXPECT().Set(ctx, gomock.Any()).Return(nil)

		err := mockSession.Set(ctx, &mockRecord{})
		assert.NoError(t, err)
	})

	t.Run("set error", func(t *testing.T) {
		expectedErr := errors.New("set error")
		mockSession.EXPECT().Set(ctx, gomock.Any()).Return(expectedErr)

		err := mockSession.Set(ctx, &mockRecord{})
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("set multi success", func(t *testing.T) {
		records := []dal.Record{&mockRecord{}}
		mockSession.EXPECT().SetMulti(ctx, gomock.Any()).Return(nil)
		err := mockSession.SetMulti(ctx, records)
		assert.NoError(t, err)
	})

	t.Run("set multi error", func(t *testing.T) {
		records := []dal.Record{&mockRecord{}}
		expectedErr := errors.New("set multi error")
		mockSession.EXPECT().SetMulti(ctx, gomock.Any()).Return(expectedErr)
		err := mockSession.SetMulti(ctx, records)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockWriteSession_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockWriteSession(ctrl)
	ctx := context.Background()
	key := &dal.Key{}
	updates := []update.Update{}

	t.Run("update success", func(t *testing.T) {
		mockSession.EXPECT().Update(ctx, key, updates).Return(nil)

		err := mockSession.Update(ctx, key, updates)
		assert.NoError(t, err)
	})

	t.Run("update error", func(t *testing.T) {
		expectedErr := errors.New("update error")
		mockSession.EXPECT().Update(ctx, key, updates).Return(expectedErr)

		err := mockSession.Update(ctx, key, updates)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("update multi success", func(t *testing.T) {
		keys := []*dal.Key{key}
		mockSession.EXPECT().UpdateMulti(ctx, gomock.Any(), gomock.Any()).Return(nil)
		err := mockSession.UpdateMulti(ctx, keys, updates)
		assert.NoError(t, err)
	})

	t.Run("update multi error", func(t *testing.T) {
		keys := []*dal.Key{key}
		expectedErr := errors.New("update multi error")
		mockSession.EXPECT().UpdateMulti(ctx, gomock.Any(), gomock.Any()).Return(expectedErr)
		err := mockSession.UpdateMulti(ctx, keys, updates)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("update record success", func(t *testing.T) {
		record := &mockRecord{}
		mockSession.EXPECT().UpdateRecord(ctx, gomock.Any(), updates).Return(nil)
		err := mockSession.UpdateRecord(ctx, record, updates)
		assert.NoError(t, err)
	})

	t.Run("update record error", func(t *testing.T) {
		record := &mockRecord{}
		expectedErr := errors.New("update record error")
		mockSession.EXPECT().UpdateRecord(ctx, gomock.Any(), updates).Return(expectedErr)
		err := mockSession.UpdateRecord(ctx, record, updates)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestNewMockReadwriteSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadwriteSession(ctrl)
	assert.NotNil(t, mockSession)
	assert.NotNil(t, mockSession.EXPECT())
}

func TestMockReadwriteSession_Exists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := NewMockReadwriteSession(ctrl)
	ctx := context.Background()
	key := &dal.Key{}

	t.Run("exists returns true", func(t *testing.T) {
		mockSession.EXPECT().Exists(ctx, key).Return(true, nil)

		exists, err := mockSession.Exists(ctx, key)
		assert.True(t, exists)
		assert.NoError(t, err)
	})

	t.Run("exists returns error", func(t *testing.T) {
		expectedErr := errors.New("exists error")
		mockSession.EXPECT().Exists(ctx, key).Return(false, expectedErr)

		exists, err := mockSession.Exists(ctx, key)
		assert.False(t, exists)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockReadwriteSession_ReadQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := NewMockReadwriteSession(ctrl)
	ctx := context.Background()

	t.Run("get success", func(t *testing.T) {
		s.EXPECT().Get(ctx, gomock.Any()).Return(nil)
		err := s.Get(ctx, &mockRecord{})
		assert.NoError(t, err)
	})

	t.Run("get multi success", func(t *testing.T) {
		records := []dal.Record{&mockRecord{}}
		s.EXPECT().GetMulti(ctx, gomock.Any()).Return(nil)
		err := s.GetMulti(ctx, records)
		assert.NoError(t, err)
	})

	t.Run("query all success", func(t *testing.T) {
		expected := []dal.Record{&mockRecord{}}
		s.EXPECT().QueryAllRecords(ctx, gomock.Any()).Return(expected, nil)
		records, err := s.QueryAllRecords(ctx, nil)
		assert.NoError(t, err)
		assert.Equal(t, expected, records)
	})

	t.Run("query reader error", func(t *testing.T) {
		expectedErr := errors.New("reader error")
		s.EXPECT().QueryReader(ctx, gomock.Any()).Return(nil, expectedErr)
		reader, err := s.QueryReader(ctx, nil)
		assert.Error(t, err)
		assert.Nil(t, reader)
		assert.Equal(t, expectedErr, err)
	})
}

func TestMockReadwriteSession_Writes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := NewMockReadwriteSession(ctrl)
	ctx := context.Background()
	key := &dal.Key{}
	updates := []update.Update{}

	t.Run("delete/deleteMulti", func(t *testing.T) {
		s.EXPECT().Delete(ctx, key).Return(nil)
		assert.NoError(t, s.Delete(ctx, key))

		keys := []*dal.Key{key}
		s.EXPECT().DeleteMulti(ctx, gomock.Any()).Return(nil)
		assert.NoError(t, s.DeleteMulti(ctx, keys))
	})

	t.Run("insert/insertMulti", func(t *testing.T) {
		s.EXPECT().Insert(ctx, gomock.Any()).Return(nil)
		assert.NoError(t, s.Insert(ctx, &mockRecord{}))

		records := []dal.Record{&mockRecord{}}
		s.EXPECT().InsertMulti(ctx, gomock.Any()).Return(nil)
		assert.NoError(t, s.InsertMulti(ctx, records))
	})

	t.Run("set/setMulti", func(t *testing.T) {
		s.EXPECT().Set(ctx, gomock.Any()).Return(nil)
		assert.NoError(t, s.Set(ctx, &mockRecord{}))

		records := []dal.Record{&mockRecord{}}
		s.EXPECT().SetMulti(ctx, gomock.Any()).Return(nil)
		assert.NoError(t, s.SetMulti(ctx, records))
	})

	t.Run("update/updateMulti/updateRecord", func(t *testing.T) {
		s.EXPECT().Update(ctx, key, updates).Return(nil)
		assert.NoError(t, s.Update(ctx, key, updates))

		keys := []*dal.Key{key}
		s.EXPECT().UpdateMulti(ctx, gomock.Any(), gomock.Any()).Return(nil)
		assert.NoError(t, s.UpdateMulti(ctx, keys, updates))

		s.EXPECT().UpdateRecord(ctx, gomock.Any(), updates).Return(nil)
		assert.NoError(t, s.UpdateRecord(ctx, &mockRecord{}, updates))
	})
}
