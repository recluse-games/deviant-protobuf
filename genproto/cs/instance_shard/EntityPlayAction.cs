// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: entity_play_action.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from entity_play_action.proto</summary>
  public static partial class EntityPlayActionReflection {

    #region Descriptor
    /// <summary>File descriptor for entity_play_action.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static EntityPlayActionReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChhlbnRpdHlfcGxheV9hY3Rpb24ucHJvdG8SB0RldmlhbnQaCnBsYXkucHJv",
            "dG8iTAoQRW50aXR5UGxheUFjdGlvbhIKCgJpZBgBIAEoCRIOCgZjYXJkSWQY",
            "AiABKAkSHAoFcGxheXMYAyADKAsyDS5EZXZpYW50LlBsYXlCC1oJLjtkZXZp",
            "YW50YgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Deviant.PlayReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Deviant.EntityPlayAction), global::Deviant.EntityPlayAction.Parser, new[]{ "Id", "CardId", "Plays" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class EntityPlayAction : pb::IMessage<EntityPlayAction> {
    private static readonly pb::MessageParser<EntityPlayAction> _parser = new pb::MessageParser<EntityPlayAction>(() => new EntityPlayAction());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<EntityPlayAction> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Deviant.EntityPlayActionReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EntityPlayAction() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EntityPlayAction(EntityPlayAction other) : this() {
      id_ = other.id_;
      cardId_ = other.cardId_;
      plays_ = other.plays_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EntityPlayAction Clone() {
      return new EntityPlayAction(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private string id_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Id {
      get { return id_; }
      set {
        id_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "cardId" field.</summary>
    public const int CardIdFieldNumber = 2;
    private string cardId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string CardId {
      get { return cardId_; }
      set {
        cardId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "plays" field.</summary>
    public const int PlaysFieldNumber = 3;
    private static readonly pb::FieldCodec<global::Deviant.Play> _repeated_plays_codec
        = pb::FieldCodec.ForMessage(26, global::Deviant.Play.Parser);
    private readonly pbc::RepeatedField<global::Deviant.Play> plays_ = new pbc::RepeatedField<global::Deviant.Play>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Deviant.Play> Plays {
      get { return plays_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as EntityPlayAction);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(EntityPlayAction other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (CardId != other.CardId) return false;
      if(!plays_.Equals(other.plays_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id.Length != 0) hash ^= Id.GetHashCode();
      if (CardId.Length != 0) hash ^= CardId.GetHashCode();
      hash ^= plays_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Id);
      }
      if (CardId.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(CardId);
      }
      plays_.WriteTo(output, _repeated_plays_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Id);
      }
      if (CardId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(CardId);
      }
      size += plays_.CalculateSize(_repeated_plays_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(EntityPlayAction other) {
      if (other == null) {
        return;
      }
      if (other.Id.Length != 0) {
        Id = other.Id;
      }
      if (other.CardId.Length != 0) {
        CardId = other.CardId;
      }
      plays_.Add(other.plays_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Id = input.ReadString();
            break;
          }
          case 18: {
            CardId = input.ReadString();
            break;
          }
          case 26: {
            plays_.AddEntriesFrom(input, _repeated_plays_codec);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
