package gcanpb

func (m *MessageSet) CheckIntegrity() error {
	for _, m := range m.Messages {
		if err := m.CheckIntegrity(); err != nil {
			return err
		}
	}
	return nil
}
