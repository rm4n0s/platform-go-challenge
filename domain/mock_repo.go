package domain

import "context"

type MockDB struct {
	addAsset    func(ctx context.Context, asset InputAsset) (*Asset, error)
	updateAsset func(ctx context.Context, assetID uint, asset InputAsset) (*Asset, error)
	userExists  func(ctx context.Context, username string) (bool, error)
	addUser     func(ctx context.Context, user User) (*User, error)
	findUser    func(ctx context.Context, username string) (*User, error)
}

func (d *MockDB) AddAsset(ctx context.Context, asset InputAsset) (*Asset, error) {
	return d.addAsset(ctx, asset)
}
func (d *MockDB) DeleteAsset(ctx context.Context, at AssetType, assetID uint) error {
	return nil
}
func (d *MockDB) UpdateAsset(ctx context.Context, assetID uint, asset InputAsset) (*Asset, error) {
	return d.updateAsset(ctx, assetID, asset)
}
func (d *MockDB) GetAsset(ctx context.Context, at AssetType, assetID uint) (*Asset, error) {
	return nil, nil
}
func (d *MockDB) ListAssets(ctx context.Context, query QueryAssets) (*ListedAssets, error) {
	return nil, nil
}
func (d *MockDB) FavouriteAsset(ctx context.Context, userID, assetID uint, at AssetType, isFavourite bool) (uint, error) {
	return 0, nil
}
func (d *MockDB) ListFavouriteAssets(ctx context.Context, userID uint, onlyFav bool, query QueryAssets) (*ListedAssets, error) {
	return nil, nil
}

func (d *MockDB) RemoveFavouriteAssetFromEveryone(ctx context.Context, assetID uint, at AssetType) error {
	return nil
}
func (d *MockDB) AddUser(ctx context.Context, user User) (*User, error) {
	return d.addUser(ctx, user)
}
func (d *MockDB) FindUser(ctx context.Context, username string) (*User, error) {
	return d.findUser(ctx, username)
}
func (d *MockDB) UserExists(ctx context.Context, username string) (bool, error) {
	return d.userExists(ctx, username)
}
func (d *MockDB) GetUser(ctx context.Context, userID uint) (*User, error) {
	return nil, nil
}
